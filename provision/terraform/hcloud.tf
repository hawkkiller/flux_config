variable "hcloud_token" {
  sensitive = true # Requires terraform >= 0.14
}

provider "hcloud" {
    token = var.hcloud_token
}