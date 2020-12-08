# Kubernetes Auditor

[Kubernetes Auditor by AppsCode](https://github.com/kubeshield/auditor) - Kubernetes Auditor for Kubernetes

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install auditor appscode/auditor -n kube-system
```

## Introduction

This chart deploys a Grafana operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the chart with the release name `auditor`:

```console
$ helm install auditor appscode/auditor -n kube-system
```

The command deploys a Grafana operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `auditor`:

```console
$ helm delete auditor -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `auditor` chart and their default values.

|            Parameter             |                                                                                                             Description                                                                                                              |                                Default                                |
|----------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------|
| nameOverride                     | Overrides name template                                                                                                                                                                                                              | `""`                                                                  |
| fullnameOverride                 | Overrides fullname template                                                                                                                                                                                                          | `""`                                                                  |
| replicaCount                     | Number of Grafana operator replicas to create (only 1 is supported)                                                                                                                                                                  | `1`                                                                   |
| image.registry                   | Docker registry used to pull operator image                                                                                                                                                                                          | `kubeshield`                                                          |
| image.repository                 | Name of operator container image                                                                                                                                                                                                     | `auditor`                                                             |
| image.tag                        | Operator container image tag                                                                                                                                                                                                         | `v0.0.1`                                                              |
| image.resources                  | Compute Resources required by the operator container                                                                                                                                                                                 | `{}`                                                                  |
| image.securityContext            | Security options the operator container should run with                                                                                                                                                                              | `{}`                                                                  |
| imagePullSecrets                 | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/auditor \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1` | `[]`                                                                  |
| imagePullPolicy                  | Container image pull policy                                                                                                                                                                                                          | `IfNotPresent`                                                        |
| criticalAddon                    | If true, installs Grafana operator as critical addon                                                                                                                                                                                 | `false`                                                               |
| logLevel                         | Log level for operator                                                                                                                                                                                                               | `3`                                                                   |
| annotations                      | Annotations applied to operator deployment                                                                                                                                                                                           | `{}`                                                                  |
| podAnnotations                   | Annotations passed to operator pod(s).                                                                                                                                                                                               | `{}`                                                                  |
| nodeSelector                     | Node labels for pod assignment                                                                                                                                                                                                       | `{"beta.kubernetes.io/arch":"amd64","beta.kubernetes.io/os":"linux"}` |
| tolerations                      | Tolerations for pod assignment                                                                                                                                                                                                       | `[]`                                                                  |
| affinity                         | Affinity rules for pod assignment                                                                                                                                                                                                    | `{}`                                                                  |
| podSecurityContext               | Security options the operator pod should run with.                                                                                                                                                                                   | `{"fsGroup":65535}`                                                   |
| serviceAccount.create            | Specifies whether a service account should be created                                                                                                                                                                                | `true`                                                                |
| serviceAccount.annotations       | Annotations to add to the service account                                                                                                                                                                                            | `{}`                                                                  |
| serviceAccount.name              | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                               | ``                                                                    |
| enableAnalytics                  | If true, sends usage analytics                                                                                                                                                                                                       | `true`                                                                |
| monitoring.agent                 | Name of monitoring agent (either "prometheus.io/operator" or "prometheus.io/builtin")                                                                                                                                                | `"none"`                                                              |
| monitoring.server                | Specify whether to monitor Grafana operator                                                                                                                                                                                          | `false`                                                               |
| monitoring.prometheus.namespace  | Specify the namespace where Prometheus server is running or will be deployed.                                                                                                                                                        | `""`                                                                  |
| monitoring.serviceMonitor.labels | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                  | `{}`                                                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install auditor appscode/auditor -n kube-system --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install auditor appscode/auditor -n kube-system --values values.yaml
```
