# Moloon - distributed serverless Edge

## Building

```bash
docker-compose build
```

## Kubernetes dependencies

Make sure you have the dependency set to the right version, according to [this](https://github.com/kubernetes/client-go/blob/master/INSTALL.md#go-modules). As an example for using Kubernetes 1.15 client use

```bash
go get k8s.io/client-go@kubernetes-1.15.0
```

