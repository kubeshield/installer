# Identity Server

[Identity Server by AppsCode](https://github.com/kubeshield/identity-server)

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install identity-server appscode/identity-server -n kube-system
```

## Introduction

This chart bootstraps an [Identity Server](https://github.com/kubeshield/identity-server) deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the chart with the release name `identity-server`:

```console
$ helm install identity-server appscode/identity-server -n kube-system
```

The command deploys Identity Server on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `identity-server`:

```console
$ helm uninstall identity-server -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the Identity Server and their default values.

| Parameter                               | Description                                                                                                                                                                                                                                                                                                                                                 | Default                                                   |
| --------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------- |
| `replicaCount`                          | Number of Identity Server replicas to create (only 1 is supported)                                                                                                                                                                                                                                                                                          | `1`                                                       |
| `image.registry`                        | Docker registry used to pull Identity Server image                                                                                                                                                                                                                                                                                                          | `kubeshield`                                              |
| `image.repository`                      | Identity Server container image                                                                                                                                                                                                                                                                                                                             | `identity-server`                                         |
| `image.tag`                             | Identity Server container image tag                                                                                                                                                                                                                                                                                                                         | `v0.1.0`                                                  |
| `imagePullSecrets`                      | Specify image pull secrets                                                                                                                                                                                                                                                                                                                                  | `[]`                                                      |
| `imagePullPolicy`                       | Image pull policy                                                                                                                                                                                                                                                                                                                                           | `IfNotPresent`                                            |
| `criticalAddon`                         | If true, installs Identity Server as critical addon                                                                                                                                                                                                                                                                                                         | `false`                                                   |
| `logLevel`                              | Log level for operator                                                                                                                                                                                                                                                                                                                                      | `3`                                                       |
| `affinity`                              | Affinity rules for pod assignment                                                                                                                                                                                                                                                                                                                           | `{}`                                                      |
| `annotations`                           | Annotations applied to operator deployment                                                                                                                                 | `{}`                                                      |
| `podAnnotations`                        | Annotations applied to operator pod(s)                                                                                                                                     | `{}`                                                      |
| `nodeSelector`                          | Node labels for pod assignment                                                                                                                                                                                                                                                                                                                              | `{}`                                                      |
| `tolerations`                           | Tolerations used pod assignment                                                                                                                                                                                                                                                                                                                             | `{}`                                                      |
| `serviceAccount.create`                 | If `true`, create a new service account                                                                                                                                                                                                                                                                                                                     | `true`                                                    |
| `serviceAccount.name`                   | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template                                                                                                                                                                                                                               | ``                                                        |
| `apiserver.groupPriorityMinimum`        | The minimum priority the group should have.                                                                                                                                                                                                                                                                                                                 | 10000                                                     |
| `apiserver.versionPriority`             | The ordering of this API inside of the group.                                                                                                                                                                                                                                                                                                               | 15                                                        |
| `apiserver.useKubeapiserverFqdnForAks`  | If true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522                                                                                                                                                                                                                                                     | `true`                                                    |
| `apiserver.healthcheck.enabled`         | Enable readiness and liveliness probes                                                                                                                                                                                                                                                                                                                      | `false`                                                   |
| `apiserver.servingCerts.generate`       | If true, generate on install/upgrade the certs that allow the kube-apiserver (and potentially ServiceMonitor) to authenticate identity-server pods. Otherwise specify in `apiserver.servingCerts.{caCrt, serverCrt, serverKey}`. See also: [example terraform](https://github.com/kubeshield/identity-server/blob/master/charts/identity-server/example-terraform.tf) | `true`                                          |
| `enableAnalytics`                       | Send usage events to Google Analytics                                                                                                                                                                                                                                                                                                                       | `true`                                                    |
| `monitoring.agent`                      | Specify which monitoring agent to use for monitoring Identity Server. It accepts either `prometheus.io/builtin` or `prometheus.io/operator`.                                                                                                                                                                                                         | `none`                                                    |
| `monitoring.operator`                   | Specify whether to monitor Identity Server.                                                                                                                                                                                                                                                                                                                 | `false`                                                   |
| `monitoring.prometheus.namespace`       | Specify the namespace where Prometheus server is running or will be deployed.                                                                                                                                                                                                                                                                               | Release namespace                                         |
| `monitoring.serviceMonitor.labels`      | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                                                                                                                                  | `app: <generated app name>` and `release: <release name>` |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install identity-server appscode/identity-server -n kube-system --set image.tag=v0.3.0
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install identity-server appscode/identity-server -n kube-system --values values.yaml
```
