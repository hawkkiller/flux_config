# TF setup

terraform {
  required_providers {
    hcloud = {
      source  = "hetznercloud/hcloud"
      version = "1.35.2"
    }
    talos = {
      source  = "siderolabs/talos"
      version = "0.2.0"
    }
  }
}
