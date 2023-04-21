# Flux configs

This repository contains a set of Flux configs for deploying various services to a Kubernetes cluster.

## Useful commands

- `flux get kustomizations` - list all the kustomizations in the cluster
- `flux get helmreleases` - list all the Helm releases in the cluster

## DNS01 solver setup

- `gcloud iam service-accounts create dns01-solver --display-name "dns01-solver"`
- `gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member serviceAccount:dns01-solver@$PROJECT_ID.iam.gserviceaccount.com \
  --role roles/dns.admin`