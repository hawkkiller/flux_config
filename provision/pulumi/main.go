package main

import (
	"fmt"

	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/siderolabs/pulumi-provider-talos/sdk/go/talos"
)

type Resources struct {
	Config                                *PulumiConfig
	Network                               *hcloud.Network
	NetworkSubnet                         *hcloud.NetworkSubnet
	LoadBalancer                          *hcloud.LoadBalancer
	ControlPlanes                         []*hcloud.Server
	TalosMachineSecrets                   *talos.TalosMachineSecrets
	TalosClientConfiguration              *talos.TalosClientConfiguration
	TalosControlPlaneMachineConfiguration *talos.TalosMachineConfigurationControlplane
	TalosWorkerMachineConfiguration       *talos.TalosMachineConfigurationWorker
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		config, err := ParseConfig()

		if err != nil {
			return err
		}

		fmt.Printf("Using config: %+v\n", config)

		resources := Resources{
			Config: config,
		}

		err = resources.CreateNetworks(ctx)
		if err != nil {
			return err
		}

		err = resources.CreateSubnets(ctx)

		if err != nil {
			return err
		}

		err = resources.CreateLoadBalancer(ctx)

		if err != nil {
			return err
		}

		err = resources.CreateLoadBalancerNetwork(ctx)

		if err != nil {
			return err
		}

		err = resources.CreateControlPlanes(ctx)

		if err != nil {
			return err
		}

		err = resources.CreateLoadBalancerTargets(ctx)

		if err != nil {
			return err
		}

		err = resources.CreateLoadBalancerServices(ctx)

		if err != nil {
			return err
		}

		return nil
	})
}
