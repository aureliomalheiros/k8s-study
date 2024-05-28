# Helm configuration

```markdown
helm-cat/
├── charts/
│   ├── templates/
│       ├── namespace.yaml
│       ├── deployment.yaml
│       ├── service.yaml
│       ├── serviceaccount.yaml
│       ├── clusterrole.yaml
│       ├── clusterrolebinding.yaml
│       ├── daemonset.yaml
│       ├── secret.yaml
│       ├── tests/
│           ├── test-connection.yaml
│   ├── Chart.yaml
│   ├── README.md
│   ├── _helpers.tpl
├── values/
│   ├── k8s-monitoring/
│       ├── newrelic.yaml
```
