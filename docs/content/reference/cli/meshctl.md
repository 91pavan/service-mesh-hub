---
title: "meshctl"
weight: 5
---
## meshctl

CLI for Service Mesh Hub

### Synopsis

CLI for Service Mesh Hub

### Options

```
      --context string          Specify which context from the kubeconfig should be used; uses current context if none is specified
  -h, --help                    help for meshctl
      --kube-timeout duration   Specify the default timeout for requests to kubernetes API servers. (default 5s)
      --kubeconfig string       Specify the kubeconfig for the current command
  -n, --namespace string        Specify the namespace where Service Mesh Hub resources should be written (default "service-mesh-hub")
  -v, --verbose                 Enable verbose mode, which outputs additional execution details that may be helpful for debugging
```

### SEE ALSO

* [meshctl check](../meshctl_check)	 - Check the status of a Service Mesh Hub installation
* [meshctl cluster](../meshctl_cluster)	 - Register and perform operations on clusters
* [meshctl install](../meshctl_install)	 - Install Service Mesh Hub
* [meshctl istio](../meshctl_istio)	 - Manage installations of Istio
* [meshctl uninstall](../meshctl_uninstall)	 - Completely uninstall Service Mesh Hub and remove associated CRDs
* [meshctl upgrade](../meshctl_upgrade)	 - In-place upgrade of the meshctl binary
* [meshctl version](../meshctl_version)	 - Display the version of meshctl and Service Mesh Hub server components

