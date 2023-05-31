package main

// func (r *Resources) CreateLoadBalancer(ctx *pulumi.Context) error {
// 	lb, err := hcloud.NewLoadBalancer(ctx, r.Config.Network.LoadBalancer.Name, &hcloud.LoadBalancerArgs{
// 		Algorithm: hcloud.LoadBalancerAlgorithmArgs{
// 			Type: pulumi.String(r.Config.Network.LoadBalancer.Algorithm),
// 		},
// 		Labels:           pulumi.Map{"type": pulumi.String("cluster")},
// 		LoadBalancerType: pulumi.String(r.Config.Network.LoadBalancer.Type),
// 		NetworkZone:      pulumi.String(r.Config.Network.Zone),
// 		Name:             pulumi.String(r.Config.Network.LoadBalancer.Name),
// 	}, pulumi.DependsOn([]pulumi.Resource{r.Network}))

// 	if err != nil {
// 		return errors.Wrap(err, "failed to create load balancer")
// 	}

// 	r.LoadBalancer = lb
// 	return nil
// }
