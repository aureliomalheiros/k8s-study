# Lab study Kubernetes and Service Mesh

![Don't panic](https://img.shields.io/badge/env-don't%20panic-green?style=for-the-badge&logo=appveyor)
## Requirements

- [Docker](https://docs.docker.com/engine/install/)
- [wget](https://www.tecmint.com/install-wget-in-linux/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [k3d](https://k3d.io/v5.4.6/)

### Create cluster

**K3D**

Execute command:

```bash
k3d cluster create name_cluster --servers number_nodes
```

Example:
```bash
k3d cluster create tribix --servers 4
```

Destroy cluster

```bash
k3d cluster delete name_cluster
```

**MINIKUBE**

[Start Minikube](https://minikube.sigs.k8s.io/docs/start/)

`minikube start --nodes=3 --driver=virtualbox --cpus=2 --memory=3000`


### ISTIO

[ISTIO Install](https://istio.io/latest/docs/setup/install/istioctl/)

