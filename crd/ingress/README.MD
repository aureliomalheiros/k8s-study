# Ingress controller
![Build](https://img.shields.io/badge/build-passing-green?style=for-the-badge)
## Requirements

- Install ingress nginx
    - [Ingress Nginx - yaml](https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.4.0/deploy/static/provider/baremetal/deploy.yaml)
    - [Ingress Nginx - Doc](https://kubernetes.github.io/ingress-nginx/deploy/#bare-metal-clusters)

**Create cluster using the command below:**

> **Note**
> The parameter `-p` similar `port-forwarding` of the kubectl.

```bash
k3d cluster create --servers 3 -p "80:30000@loadbalancer"
```

