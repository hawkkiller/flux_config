package main

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func (r *Resources) CreateControlPlanes(ctx *pulumi.Context) error {
	for _, cp := range r.Config.ControlPlanes {
		plane, err := hcloud.NewServer(ctx, cp.Name, &hcloud.ServerArgs{
			Name:       pulumi.String(cp.Name),
			ServerType: pulumi.String(cp.Type),
			Image:      pulumi.String(pulumi.String(r.Config.ImageID)),
			Location:   pulumi.String(cp.Location),
			Networks: hcloud.ServerNetworkTypeArray{
				hcloud.ServerNetworkTypeArgs{
					NetworkId: r.Network.ID().ApplyT(strconv.Atoi).(pulumi.IntOutput),
					Ip:        pulumi.String(cp.Ip),
				},
			},
			Labels: pulumi.Map{},
		}, pulumi.DependsOn([]pulumi.Resource{r.Network, r.LoadBalancer}))

		if err != nil {
			return errors.Wrap(err, "failed to create control plane")
		}

		r.ControlPlanes = append(r.ControlPlanes, plane)
	}

	return nil
}
