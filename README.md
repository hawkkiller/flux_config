# LazebnyIO

The project is a set of configs to deploy various self-hosted
applications to Kubernetes Cluster using FluxCD.

## Prerequisites

- Kubernetes Cluster
- Sops installed
- FluxCD installed (see [FluxCD docs](https://fluxcd.io/docs/get-started/))

## Installation

1. `export GITHUB_TOKEN=<your github token>`
2. kubectl create ns flux-system --dry-run=client -o yaml | kubectl apply -f -
3. cat age.agekey |
kubectl create secret generic sops-age \
--namespace=flux-system \
--from-file=age.agekey=/dev/stdin
4. flux bootstrap github \
  --token-auth \
  --owner=hawkkiller \
  --repository=flux_config \
  --branch=main \
  --path=./kubernetes/flux \
  --components-extra=image-reflector-controller,image-automation-controller \
  --version=latest \
  --personal
