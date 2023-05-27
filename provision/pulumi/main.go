package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/siderolabs/pulumi-provider-talos/sdk/go/talos"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		config, err := ParseConfig()

		if err != nil {
			return err
		}

		fmt.Printf("Using config: %+v\n", config)

		clusterNetwork, err := hcloud.NewNetwork(ctx, config.Network.Name, &hcloud.NetworkArgs{
			IpRange: pulumi.String(config.Network.IpRange),
			Labels: pulumi.Map{
				"type": pulumi.String("cluster"),
			},
		})

		if err != nil {
			return errors.Wrap(err, "failed to create network")
		}

		clusterNetworkID := clusterNetwork.ID().ApplyT(strconv.Atoi).(pulumi.IntOutput)

		cpSubnet, err := hcloud.NewNetworkSubnet(ctx, "cp-subnet", &hcloud.NetworkSubnetArgs{
			IpRange:     pulumi.String("10.0.2.0/24"),
			NetworkId:   clusterNetworkID,
			NetworkZone: pulumi.String(config.Network.Zone),
			Type:        pulumi.String("cloud"),
		}, pulumi.DependsOn([]pulumi.Resource{clusterNetwork}))

		if err != nil {
			return errors.Wrap(err, "failed to create subnet")
		}

		workerSubnet, err := hcloud.NewNetworkSubnet(ctx, "worker-subnet", &hcloud.NetworkSubnetArgs{
			IpRange:     pulumi.String("10.0.1.0/24"),
			NetworkId:   clusterNetworkID,
			NetworkZone: pulumi.String(config.Network.Zone),
			Type:        pulumi.String("cloud"),
		}, pulumi.DependsOn([]pulumi.Resource{clusterNetwork}))

		if err != nil {
			return errors.Wrap(err, "failed to create subnet")
		}

		lb, err := hcloud.NewLoadBalancer(ctx, config.Network.LoadBalancer.Name, &hcloud.LoadBalancerArgs{
			Algorithm: hcloud.LoadBalancerAlgorithmArgs{
				Type: pulumi.String(config.Network.LoadBalancer.Algorithm),
			},
			Labels:           pulumi.Map{"type": pulumi.String("cluster")},
			LoadBalancerType: pulumi.String(config.Network.LoadBalancer.Type),
			NetworkZone:      pulumi.String(config.Network.Zone),
			Name:             pulumi.String(config.Network.LoadBalancer.Name),
		})

		if err != nil {
			return errors.Wrap(err, "failed to create load balancer")
		}

		_, err = hcloud.NewLoadBalancerNetwork(ctx, config.Network.LoadBalancer.Name, &hcloud.LoadBalancerNetworkArgs{
			// NetworkId:      r.Network.ID().ApplyT(strconv.Atoi).(pulumi.IntOutput),
			LoadBalancerId: lb.ID().ApplyT(strconv.Atoi).(pulumi.IntOutput),
			SubnetId:       cpSubnet.ID(),
		}, pulumi.DependsOn([]pulumi.Resource{cpSubnet, lb}))

		if err != nil {
			return errors.Wrap(err, "failed to create load balancer network")
		}

		cpsIpIpv4 := make([]pulumi.StringInput, len(config.ControlPlanes))

		for i, cp := range config.ControlPlanes {
			plane, err := hcloud.NewServer(ctx, cp.Name, &hcloud.ServerArgs{
				Name:       pulumi.String(cp.Name),
				ServerType: pulumi.String(cp.Type),
				Image:      pulumi.String(pulumi.String(config.ImageID)),
				Location:   pulumi.String(cp.Location),
				Networks: hcloud.ServerNetworkTypeArray{
					hcloud.ServerNetworkTypeArgs{
						NetworkId: cpSubnet.NetworkId,
					},
				},
				Labels: pulumi.Map{
					"type": pulumi.String("control-plane"),
				},
			}, pulumi.DependsOn([]pulumi.Resource{cpSubnet}))

			if err != nil {
				return errors.Wrap(err, "failed to create control plane")
			}
			cpsIpIpv4[i] = plane.Ipv4Address

			if err != nil {
				return errors.Wrap(err, "failed to create load balancer target")
			}
		}

		// Create LB target
		_, err = hcloud.NewLoadBalancerTarget(ctx, "control-planes", &hcloud.LoadBalancerTargetArgs{
			LoadBalancerId: lb.ID().ApplyT(strconv.Atoi).(pulumi.IntOutput),
			Type:           pulumi.String("label_selector"),
			LabelSelector:  pulumi.String("type=control-plane"),
		}, pulumi.DependsOn([]pulumi.Resource{lb}))

		if err != nil {
			return errors.Wrap(err, "failed to create load balancer target")
		}

		// loadbalance kubectl traffic
		lbKubectl, err := hcloud.NewLoadBalancerService(ctx, "control-plane-kubectl", &hcloud.LoadBalancerServiceArgs{
			LoadBalancerId:  lb.ID(),
			Protocol:        pulumi.String("tcp"),
			ListenPort:      pulumi.Int(6443),
			DestinationPort: pulumi.Int(6443),
		}, pulumi.DependsOn([]pulumi.Resource{lb}))

		if err != nil {
			return errors.Wrap(err, "failed to create load balancer service")
		}

		// loadbalance talosctl traffic
		lbTalosctl, err := hcloud.NewLoadBalancerService(ctx, "control-plane-talosctl", &hcloud.LoadBalancerServiceArgs{
			LoadBalancerId:  lb.ID(),
			Protocol:        pulumi.String("tcp"),
			ListenPort:      pulumi.Int(50000),
			DestinationPort: pulumi.Int(50000),
		}, pulumi.DependsOn([]pulumi.Resource{lb}))

		if err != nil {
			return errors.Wrap(err, "failed to create load balancer service")
		}

		talosMachineSecrets, err := talos.NewTalosMachineSecrets(ctx, "talos-machine-secrets", nil)

		if err != nil {
			return errors.Wrap(err, "failed to create talos machine secrets")
		}

		talosCfg, err := talos.NewTalosClientConfiguration(ctx, "talos-client-configuration", &talos.TalosClientConfigurationArgs{
			ClusterName:    pulumi.String(config.ClusterName),
			MachineSecrets: talosMachineSecrets.MachineSecrets,
			Endpoints:      pulumi.StringArray(cpsIpIpv4),
			// Nodes:          pulumi.StringArray{cpsIpIpv4[0]},
		}, pulumi.DependsOn([]pulumi.Resource{talosMachineSecrets, lb, lbKubectl, lbTalosctl}))

		if err != nil {
			return errors.Wrap(err, "failed to create talos client configuration")
		}

		talosCpMachineConfig, err := talos.NewTalosMachineConfigurationControlplane(
			ctx, "talos-control-plane-machine-configuration",
			&talos.TalosMachineConfigurationControlplaneArgs{
				ClusterName:     talosCfg.ClusterName,
				ClusterEndpoint: pulumi.Sprintf("https://%v:6443", lb.Ipv4),
				MachineSecrets:  talosMachineSecrets.MachineSecrets,
			}, pulumi.DependsOn([]pulumi.Resource{talosCfg}),
		)

		if err != nil {
			return errors.Wrap(err, "failed to create talos control plane machine configuration")
		}

		talosWorkerMachineConfig, err := talos.NewTalosMachineConfigurationWorker(
			ctx, "talos-worker-machine-configuration",
			&talos.TalosMachineConfigurationWorkerArgs{
				ClusterName:     talosCfg.ClusterName,
				ClusterEndpoint: pulumi.Sprintf("https://%v:6443", lb.Ipv4),
				MachineSecrets:  talosMachineSecrets.MachineSecrets,
			}, pulumi.DependsOn([]pulumi.Resource{talosCfg}),
		)

		if err != nil {
			return errors.Wrap(err, "failed to create talos worker machine configuration")
		}

		for i, cpIpv4 := range cpsIpIpv4 {
			_, err := talos.NewTalosMachineConfigurationApply(ctx, fmt.Sprintf("cp-config-apply-%d", i),
				&talos.TalosMachineConfigurationApplyArgs{
					TalosConfig:          talosCfg.TalosConfig,
					MachineConfiguration: talosCpMachineConfig.MachineSecrets,
					Endpoint:             cpIpv4,
					Node:                 cpIpv4,
				}, pulumi.DependsOn([]pulumi.Resource{talosCpMachineConfig}),
			)

			if err != nil {
				return errors.Wrap(err, "failed to apply talos machine configuration")
			}
		}

		workersIpIpv4 := make([]pulumi.StringInput, len(config.Workers))

		for i, node := range config.Workers {
			worker, err := hcloud.NewServer(ctx, node.Name, &hcloud.ServerArgs{
				Name:       pulumi.String(node.Name),
				ServerType: pulumi.String(node.Type),
				Image:      pulumi.String(pulumi.String(config.ImageID)),
				Location:   pulumi.String(node.Location),
				Networks: hcloud.ServerNetworkTypeArray{
					hcloud.ServerNetworkTypeArgs{
						NetworkId: workerSubnet.NetworkId,
					},
				},
				Labels: pulumi.Map{},
			}, pulumi.DependsOn([]pulumi.Resource{workerSubnet, lb}))

			if err != nil {
				return errors.Wrap(err, "failed to create worker node")
			}

			workersIpIpv4[i] = worker.Ipv4Address
		}

		for i, workerIpv4 := range workersIpIpv4 {
			_, err := talos.NewTalosMachineConfigurationApply(ctx, fmt.Sprintf("worker-config-apply-%d", i),
				&talos.TalosMachineConfigurationApplyArgs{
					TalosConfig:          talosCfg.TalosConfig,
					MachineConfiguration: talosWorkerMachineConfig.MachineSecrets,
					Endpoint:             workerIpv4,
					Node:                 workerIpv4,
				}, pulumi.DependsOn([]pulumi.Resource{talosWorkerMachineConfig}),
			)

			if err != nil {
				return errors.Wrap(err, "failed to apply talos machine configuration")
			}
		}

		_, err = talos.NewTalosMachineBootstrap(ctx, "cp-bootstrap",
			&talos.TalosMachineBootstrapArgs{
				TalosConfig: talosCfg.TalosConfig,
				Endpoint:    cpsIpIpv4[0],
				Node:        cpsIpIpv4[0],
			}, pulumi.DependsOn([]pulumi.Resource{talosCfg}),
		)

		if err != nil {
			return errors.Wrap(err, "failed to bootstrap talos machine")
		}

		return nil
	})
}
