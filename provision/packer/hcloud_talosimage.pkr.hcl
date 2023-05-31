# hcloud.pkr.hcl

packer {
  required_plugins {
    hcloud = {
      version = ">= 1.0.0"
      source  = "github.com/hashicorp/hcloud"
    }
  }
}

variable "arch" {
  type    = string
  default = "arm64"
}

variable "talos_version" {
  type    = string
  default = "v1.4.4"
}

locals {
  image = "https://github.com/siderolabs/talos/releases/download/${var.talos_version}/hcloud-${var.arch}.raw.xz"
}

source "hcloud" "talos" {
  rescue       = "linux64"
  image        = "debian-11"
  location     = "fsn1"
  server_type  = "cax11"
  ssh_username = "root"

  snapshot_name = "talos system disk ${var.talos_version}"
  snapshot_labels = {
    type    = "infra",
    os      = "talos",
    version = "${var.talos_version}",
    arch    = "${var.arch}",
  }
}

build {
  sources = ["source.hcloud.talos"]

  provisioner "shell" {
    inline = [
      "apt-get install -y wget",
      "wget -O /tmp/talos.raw.xz ${local.image}",
      "xz -d -c /tmp/talos.raw.xz | dd of=/dev/sda && sync",
    ]
  }
}