# LazebnyIO

The project is a set of configs to deploy various self-hosted
applications to Kubernetes Cluster using FluxCD.

## Prerequisites

- Kubernetes Cluster
- Sops CLI with `age` installed
- FluxCD CLI installed

## Configuration

Steps that should be performed before installation.

### Secrets

FluxCD supports Mozilla SOPS for secrets encryption. This project uses
`age` format. To generate a new key pair use the following command:

```bash
age-keygen -o age.agekey
```

It puts both private and public keys into `age.agekey` file. The
private key should be stored in a secure place. The public key should be
saved in .sops.yaml file in the project root in "age" section.

```yaml
creation_rules:
  - path_regex: .*.ya?ml
    encrypted_regex: ^(data|stringData)$
    age: Paste it here!
```

It is needed to put the private key into kubernetes cluster as a secret. To do so, run the following command:

```bash
cat age.agekey |
kubectl create secret generic sops-age \
--namespace=flux-system \
--from-file=age.agekey=/dev/stdin
```

To create a new secret, create a secret manifest:

```yaml
apiVersion: v1
kind: Secret
metadata:
    name: secret-name
    namespace: namespace
type: Opaque
data:
  KEY: BASE64_ENCODED_VALUE
```

Then encode this file using sops:

```bash
sops -e -i secret.yaml
```

This way, this secret will be encrypted and can be stored in the git
repository. FluxCD will decrypt it during the deployment. To decrypt
the secret, use the following command:

```bash
sops -i -d secret.yaml
```

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

## Applications

The list of applications that are deployed to the cluster.

### Weave GitOps

Weave GitOps is a program that is used to track deployed
applications, sources and other Flux components.

Secret needed to deploy Weave GitOps:

```yaml
# oidc-auth.sops.yaml
apiVersion: v1
kind: Secret
metadata:
    name: oidc-auth
    namespace: flux-system
type: Opaque
data:
    # The URL of the issuer, typically the discovery URL without a path
    issuerURL: aHR0cHM6Ly9kZXgubGF6ZWJueS5pbw==
    # The client ID that has been setup for Weave GitOps in the issuer (DEX)
    clientID: BASE64_ENCODED_CLIENT_ID
    # The client secret that has been setup for Weave GitOps in the issuer (DEX)
    clientSecret: BASE64_ENCODED_CLIENT_SECRET
    # The redirect URL that has been setup for Weave GitOps in the issuer, typically the dashboard URL followed by /oauth2/callback
    redirectURL: aHR0cHM6Ly93ZWF2ZS5sYXplYm55LmlvL29hdXRoMi9jYWxsYmFjaw==
```

### Dex

Dex is an OpenID Connect provider that is used to authenticate users
in Weave GitOps.

Secret needed to deploy Dex:

```yaml
# github-client.sops.yaml
apiVersion: v1
kind: Secret
metadata:
    name: github-client
type: Opaque
data:
    # To get these values, create a new OAuth app in GitHub and use the client ID and secret
    # Note, that it is not weave gitops client id and secret
    # GITHUB_CLIENT_ID
    client-id: BASE64_ENCODED_GITHUB_CLIENT_ID
    # GITHUB_CLIENT_SECRET
    client-secret: BASE64_ENCODED_GITHUB_CLIENT_SECRET
```

You also need to configure dex/app/helmrelease.yaml to use the correct
client id and secret and redirect url.
