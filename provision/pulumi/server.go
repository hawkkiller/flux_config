package main

// func (r *Resources) CreateControlPlanes(ctx *pulumi.Context) error {
// 	for _, cp := range r.Config.ControlPlanes {
// 		plane, err := hcloud.NewServer(ctx, cp.Name, &hcloud.ServerArgs{
// 			Name:       pulumi.String(cp.Name),
// 			ServerType: pulumi.String(cp.Type),
// 			Image:      pulumi.String(pulumi.String(r.Config.ImageID)),
// 			Location:   pulumi.String(cp.Location),
// 			Networks: hcloud.ServerNetworkTypeArray{
// 				hcloud.ServerNetworkTypeArgs{
// 					NetworkId: r.Network.ID().ApplyT(strconv.Atoi).(pulumi.IntOutput),
// 				},
// 			},
// 			Labels: pulumi.Map{},
// 		}, pulumi.DependsOn([]pulumi.Resource{r.Network, r.LoadBalancer}))

// 		if err != nil {
// 			return errors.Wrap(err, "failed to create control plane")
// 		}

// 		r.ControlPlanes = append(r.ControlPlanes, plane)
// 	}

// 	return nil
// }

// func (r *Resources) CreateWorkers(ctx *pulumi.Context) error {
// 	for _, node := range r.Config.Workers {
// 		worker, err := hcloud.NewServer(ctx, node.Name, &hcloud.ServerArgs{
// 			Name:       pulumi.String(node.Name),
// 			ServerType: pulumi.String(node.Type),
// 			Image:      pulumi.String(pulumi.String(r.Config.ImageID)),
// 			Location:   pulumi.String(node.Location),
// 			Networks: hcloud.ServerNetworkTypeArray{
// 				hcloud.ServerNetworkTypeArgs{
// 					NetworkId: r.Network.ID().ApplyT(strconv.Atoi).(pulumi.IntOutput),
// 				},
// 			},
// 			Labels: pulumi.Map{},
// 		}, pulumi.DependsOn([]pulumi.Resource{r.Network, r.LoadBalancer}))

// 		if err != nil {
// 			return errors.Wrap(err, "failed to create worker node")
// 		}

// 		r.Workers = append(r.Workers, worker)
// 	}

// 	return nil
// }
