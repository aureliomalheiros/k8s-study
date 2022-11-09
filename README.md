# Lab study K3d

![Don't panic](https://img.shields.io/badge/env-don't%20panic-green?style=for-the-badge&logo=appveyor)
## Requirements

- [docker](https://docs.docker.com/engine/install/)
- [wget](https://www.tecmint.com/install-wget-in-linux/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [k3d](https://k3d.io/v5.4.6/)

### Create cluster


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
