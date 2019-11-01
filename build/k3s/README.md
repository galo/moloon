# K3s Enviroment

This document goes over teh setps of setting up a K3s envioment for Moloon, this document will setup a cluster with PC(amd64) and RPI(armv7) devices, it is based [this](https://github.com/galo/k3demo)


## Setup the Cloud system

If you are running microK8s make sure you stop the service beforfe moving forward. 

```bash
microk8s stop
```

Download an install K3s, there is many configurtation options, for a detailed setup guide follow [K3s Install Guide](https://rancher.com/docs/k3s/latest/en/quick-start/), the next steps are a sinplified cookbok.

```bash
sudo k3s server &
# Kubeconfig is written to /etc/rancher/k3s/k3s.yaml
sudo k3s kubectl get nodes
```

Copy /etc/rancher/k3s/k3s.yaml on your machine located outside the cluster as ~/.kube/config. 

```shell
sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
sudo chown galo:galo ~/.kube/config
```

You should now be ready to use kubectl commands.

```bash
kubectl get nodes
NAME    STATUS   ROLES    AGE     VERSION
horus   Ready    master   5m53s   v1.16.2-k3s.1
```

## Setup the Edge part

In this step we will add an edge node, that will use a RPI device. follow [this guide](https://github.com/galo/k3demo).

```bash
wget https://github.com/rancher/k3s/releases/download/v0.10.1/k3s-armhf && \ 
  chmod +x k3s-armhf && \ 
  sudo mv k3s-armhf /usr/local/bin/k3s
```

K3S_TOKEN is created at /var/lib/rancher/k3s/server/node-token on your server. To install on worker nodes we should pass K3S_URL along with K3S_TOKEN or K3S_CLUSTER_SECRET environment variables

```bash
sudo cat /var/lib/rancher/k3s/server/node-token
```

On the RPI device set the env variables

```bash
export K3S_TOKEN="node-tokenhere"
# For some reason DNS will not work, us eteh actual master Ip
export K3S_URL="https://192.168.0.10:6443"
```

Then startthe K3s agen on the RPI device

```bash
sudo -E k3s agent -s ${K3S_URL} -t ${K3S_TOKEN}
```

