# Prometheus Operator

I've been studying operators, and I'm describing the many uses of the Kube Prometheus oeprator for monitoring microservices, Kubernetes, etc.

## Getting Started

This is project is LAB based in the documentation [Kube-Operator](https://github.com/prometheus-operator/kube-prometheus).


```bash
minikube --nodes=2 --cpus=4 --memory=4000
kubectl create -f setup && kubectl apply -f manifests
```


