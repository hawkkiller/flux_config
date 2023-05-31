package main

// func (r *Resources) CreateNetworks(ctx *pulumi.Context) error {
// 	clusterNetwork, err := hcloud.NewNetwork(ctx, r.Config.Network.Name, &hcloud.NetworkArgs{
// 		IpRange: pulumi.String(r.Config.Network.IpRange),
// 		Labels: pulumi.Map{
// 			"type": pulumi.String("cluster"),
// 		},
// 	})

// 	if err != nil {
// 		return errors.Wrap(err, "failed to create network")
// 	}

// 	r.Network = clusterNetwork

// 	return nil
// }

// func (r *Resources) CreateSubnets(ctx *pulumi.Context) error {
// 	subnet, err := hcloud.NewNetworkSubnet(ctx, r.Config.Network.Subnet.Name, &hcloud.NetworkSubnetArgs{
// 		IpRange:     pulumi.String(r.Config.Network.Subnet.IpRange),
// 		NetworkId:   r.Network.ID().ApplyT(strconv.Atoi).(pulumi.IntOutput),
// 		NetworkZone: pulumi.String(r.Config.Network.Zone),
// 		Type:        pulumi.String("cloud"),
// 	}, pulumi.DependsOn([]pulumi.Resource{r.Network}))

// 	if err != nil {
// 		return errors.Wrap(err, "failed to create subnet")
// 	}

// 	r.NetworkSubnet = subnet
// 	return nil
// }

// func (r *Resources) CreateLoadBalancerNetwork(ctx *pulumi.Context) error {
// 	_, err := hcloud.NewLoadBalancerNetwork(ctx, r.Config.Network.LoadBalancer.Name, &hcloud.LoadBalancerNetworkArgs{
// 		// NetworkId:      r.Network.ID().ApplyT(strconv.Atoi).(pulumi.IntOutput),
// 		LoadBalancerId: r.LoadBalancer.ID().ApplyT(strconv.Atoi).(pulumi.IntOutput),
// 		SubnetId:       r.NetworkSubnet.ID(),
// 	}, pulumi.DependsOn([]pulumi.Resource{r.Network, r.LoadBalancer}))

// 	if err != nil {
// 		return errors.Wrap(err, "failed to create load balancer network")
// 	}

// 	return nil
// }

// func (r *Resources) CreateLoadBalancerTargets(ctx *pulumi.Context) error {
// 	for i, cp := range r.ControlPlanes {
// 		_, err := hcloud.NewLoadBalancerTarget(ctx, fmt.Sprintf("control-plane-%d", i), &hcloud.LoadBalancerTargetArgs{
// 			LoadBalancerId: r.LoadBalancer.ID().ApplyT(strconv.Atoi).(pulumi.IntOutput),
// 			Type:           pulumi.String("server"),
// 			ServerId:       cp.ID().ApplyT(strconv.Atoi).(pulumi.IntOutput),
// 		}, pulumi.DependsOn([]pulumi.Resource{r.LoadBalancer, cp}))

// 		if err != nil {
// 			return errors.Wrap(err, "failed to create load balancer target")
// 		}
// 	}

// 	return nil
// }

// func (r *Resources) CreateLoadBalancerServices(ctx *pulumi.Context) error {
// 	// loadbalance kubectl traffic
// 	_, err := hcloud.NewLoadBalancerService(ctx, "control-plane-kubectl", &hcloud.LoadBalancerServiceArgs{
// 		LoadBalancerId:  r.LoadBalancer.ID(),
// 		Protocol:        pulumi.String("tcp"),
// 		ListenPort:      pulumi.Int(6443),
// 		DestinationPort: pulumi.Int(6443),
// 	}, pulumi.DependsOn([]pulumi.Resource{r.LoadBalancer}))

// 	if err != nil {
// 		return errors.Wrap(err, "failed to create load balancer service")
// 	}

// 	// loadbalance talosctl traffic
// 	_, err = hcloud.NewLoadBalancerService(ctx, "control-plane-talosctl", &hcloud.LoadBalancerServiceArgs{
// 		LoadBalancerId:  r.LoadBalancer.ID(),
// 		Protocol:        pulumi.String("tcp"),
// 		ListenPort:      pulumi.Int(50000),
// 		DestinationPort: pulumi.Int(50000),
// 	}, pulumi.DependsOn([]pulumi.Resource{r.LoadBalancer}))

// 	if err != nil {
// 		return errors.Wrap(err, "failed to create load balancer service")
// 	}

// 	return nil
// }
