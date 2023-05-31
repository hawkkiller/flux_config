package main

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type PulumiConfig struct {
	Network       NetworkConfig  `yaml:"network"`
	ControlPlanes []ControlPlane `yaml:"control_planes"`
	Workers       []Worker       `yaml:"workers"`
	ImageID       string         `yaml:"image_id"`
	ClusterName   string         `yaml:"cluster_name"`
}

type NetworkConfig struct {
	IpRange      string             `yaml:"ip_range"`
	Name         string             `yaml:"name"`
	Zone         string             `yaml:"zone"`
	LoadBalancer LoadBalancerConfig `yaml:"lb"`
}

type SubnetConfig struct {
	IpRange string `yaml:"ip_range"`
	Name    string `yaml:"name"`
}

type LoadBalancerConfig struct {
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	Algorithm string `yaml:"algorithm"`
	Location  string `yaml:"location"`
}

type ControlPlane struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Location string `yaml:"location"`
	// Ip       string `yaml:"ip"`
}

type Worker struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Location string `yaml:"location"`
	// Ip       string `yaml:"ip"`
}

func ParseConfig() (*PulumiConfig, error) {
	var cfg *PulumiConfig
	bytes, err := os.ReadFile("config.yaml")

	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(bytes, &cfg)

	if err != nil {
		return cfg, errors.Wrap(err, "failed to parse config")
	}

	return cfg, nil
}
