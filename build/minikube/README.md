# Setup Moloon on Minikube

## Start the dashboard

```shell
 minikube dashboard
```

## Setup an ingress

Install the ingress add-on

```shell
minikube addons enable ingress
```

## Deploy Moolon comtroller and daemonset

Deploy the controller

```shell
 kubectl apply -f ./build/minikube/controller.yaml
```

Deploy the agents

```shell
kubectl apply -f ./build/k3s/agents.yaml
```

## Access the controler

```shell
curl http://192.168.99.101:32471/ping
```

```shell
curl http://192.168.99.101:32471/api/v1/controller/agents
```
