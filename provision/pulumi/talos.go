package main

// func (r *Resources) CreateTalosMachineSecrets(ctx *pulumi.Context) error {
// 	talosMachineSecrets, err := talos.NewTalosMachineSecrets(ctx, "talos-machine-secrets", nil)

// 	if err != nil {
// 		return errors.Wrap(err, "failed to create talos machine secrets")
// 	}

// 	r.TalosMachineSecrets = talosMachineSecrets

// 	return nil
// }

// func (r *Resources) CreateTalosClientConfiguration(ctx *pulumi.Context) error {
// 	endpoints := make([]pulumi.StringInput, len(r.ControlPlanes))

// 	for i, controlPlane := range r.ControlPlanes {
// 		endpoints[i] = controlPlane.Ipv4Address
// 	}
// 	talosCfg, err := talos.NewTalosClientConfiguration(ctx, "talos-client-configuration", &talos.TalosClientConfigurationArgs{
// 		ClusterName:    pulumi.String(r.Config.ClusterName),
// 		MachineSecrets: r.TalosMachineSecrets.MachineSecrets,
// 		Endpoints:      pulumi.StringArray(endpoints),
// 	})

// 	if err != nil {
// 		return errors.Wrap(err, "failed to create talos client configuration")
// 	}

// 	r.TalosClientConfiguration = talosCfg

// 	return nil
// }

// func (r *Resources) CreateMachineConfigurations(ctx *pulumi.Context) error {

// 	talosControlPlaneMachineConfig, err := talos.NewTalosMachineConfigurationControlplane(
// 		ctx, "talos-control-plane-machine-configuration",
// 		&talos.TalosMachineConfigurationControlplaneArgs{
// 			ClusterName:     r.TalosClientConfiguration.ClusterName,
// 			ClusterEndpoint: pulumi.Sprintf("https://%v:6443", r.LoadBalancer.Ipv4),
// 			MachineSecrets:  r.TalosMachineSecrets.MachineSecrets,
// 		}, pulumi.DependsOn([]pulumi.Resource{r.TalosMachineSecrets, r.TalosClientConfiguration}),
// 	)

// 	if err != nil {
// 		return errors.Wrap(err, "failed to create talos control plane machine configuration")
// 	}

// 	talosWorkerMachineConfig, err := talos.NewTalosMachineConfigurationWorker(
// 		ctx, "talos-worker-machine-configuration",
// 		&talos.TalosMachineConfigurationWorkerArgs{
// 			ClusterName:     r.TalosClientConfiguration.ClusterName,
// 			ClusterEndpoint: pulumi.Sprintf("https://%v:6443", r.LoadBalancer.Ipv4),
// 			MachineSecrets:  r.TalosMachineSecrets.MachineSecrets,
// 		}, pulumi.DependsOn([]pulumi.Resource{r.TalosMachineSecrets, r.TalosClientConfiguration}),
// 	)

// 	if err != nil {
// 		return errors.Wrap(err, "failed to create talos worker machine configuration")
// 	}

// 	r.TalosWorkerMachineConfiguration = talosWorkerMachineConfig
// 	r.TalosControlPlaneMachineConfiguration = talosControlPlaneMachineConfig

// 	return nil
// }

// func (r *Resources) ApplyTalosMachineConfigurationCP(ctx *pulumi.Context) error {

// 	for i, controlPlane := range r.ControlPlanes {
// 		_, err := talos.NewTalosMachineConfigurationApply(ctx, fmt.Sprintf("cp-config-apply-%d", i),
// 			&talos.TalosMachineConfigurationApplyArgs{
// 				TalosConfig:          r.TalosClientConfiguration.TalosConfig,
// 				MachineConfiguration: r.TalosControlPlaneMachineConfiguration.MachineSecrets,
// 				Endpoint:             controlPlane.Ipv4Address,
// 				Node:                 controlPlane.Ipv4Address,
// 			}, pulumi.DependsOn([]pulumi.Resource{r.TalosControlPlaneMachineConfiguration, r.TalosClientConfiguration, controlPlane}),
// 		)

// 		if err != nil {
// 			return errors.Wrap(err, "failed to apply talos machine configuration")
// 		}
// 	}

// 	return nil
// }

// func (r *Resources) ApplyTalosMachineConfigurationWorker(ctx *pulumi.Context) error {

// 	for i, worker := range r.Workers {
// 		_, err := talos.NewTalosMachineConfigurationApply(ctx, fmt.Sprintf("worker-config-apply-%d", i),
// 			&talos.TalosMachineConfigurationApplyArgs{
// 				TalosConfig:          r.TalosClientConfiguration.TalosConfig,
// 				MachineConfiguration: r.TalosWorkerMachineConfiguration.MachineSecrets,
// 				Endpoint:             worker.Ipv4Address,
// 				Node:                 worker.Ipv4Address,
// 			}, pulumi.DependsOn([]pulumi.Resource{r.TalosWorkerMachineConfiguration, r.TalosClientConfiguration, worker}),
// 		)

// 		if err != nil {
// 			return errors.Wrap(err, "failed to apply talos machine configuration")
// 		}
// 	}

// 	return nil
// }
