package inputs

import (
	"github.com/prometheus/prometheus/config"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	v1 "github.com/solo-io/supergloo/pkg/api/external/prometheus/v1"
	"gopkg.in/yaml.v2"
	kubev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PrometheusConfig(name, namespace string) *v1.PrometheusConfig {
	return &v1.PrometheusConfig{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
		Prometheus: BasicPrometheusConfig,
	}
}
func PrometheusConfigMap(name, namespace string) *kubev1.ConfigMap {
	return &kubev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: map[string]string{"prometheus.yml": BasicPrometheusConfig},
	}
}

const BasicPrometheusConfig = `
# source https://raw.githubusercontent.com/prometheus/prometheus/master/documentation/examples/prometheus-kubernetes.yml
# A scrape configuration for running Prometheus on a Kubernetes cluster.
# This uses separate scrape configs for cluster components (i.e. API server, node)
# and services to allow each to use different authentication configs.
#
# Kubernetes labels will be added as Prometheus labels on metrics via the
# ` + "`" + `labelmap` + "`" + ` relabeling action.
#
# If you are using Kubernetes 1.7.2 or earlier, please take note of the comments
# for the kubernetes-cadvisor job; you will need to edit or remove this job.

# Scrape config for API servers.
#
# Kubernetes exposes API servers as endpoints to the default/kubernetes
# service so this uses ` + "`" + `endpoints` + "`" + ` role and uses relabelling to only keep
# the endpoints associated with the default/kubernetes service using the
# default named port ` + "`" + `https` + "`" + `. This works for single API server deployments as
# well as HA API server deployments.
scrape_configs:
- job_name: 'kubernetes-apiservers'

  kubernetes_sd_configs:
  - role: endpoints

  # Default to scraping over https. If required, just disable this or change to
  # ` + "`" + `http` + "`" + `.
  scheme: https

  # This TLS & bearer token file config is used to connect to the actual scrape
  # endpoints for cluster components. This is separate to discovery auth
  # configuration because discovery & scraping are two separate concerns in
  # Prometheus. The discovery auth config is automatic if Prometheus runs inside
  # the cluster. Otherwise, more config options have to be provided within the
  # <kubernetes_sd_config>.
  tls_config:
    ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
    # If your node certificates are self-signed or use a different CA to the
    # master CA, then disable certificate verification below. Note that
    # certificate verification is an integral part of a secure infrastructure
    # so this should only be disabled in a controlled environment. You can
    # disable certificate verification by uncommenting the line below.
    #
    # insecure_skip_verify: true
  bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token

  # Keep only the default/kubernetes service endpoints for the https port. This
  # will add targets for each API server which Kubernetes adds an endpoint to
  # the default/kubernetes service.
  relabel_configs:
  - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
    action: keep
    regex: default;kubernetes;https

# Scrape config for nodes (kubelet).
#
# Rather than connecting directly to the node, the scrape is proxied though the
# Kubernetes apiserver.  This means it will work if Prometheus is running out of
# cluster, or can't connect to nodes for some other reason (e.g. because of
# firewalling).
- job_name: 'kubernetes-nodes'

  # Default to scraping over https. If required, just disable this or change to
  # ` + "`" + `http` + "`" + `.
  scheme: https

  # This TLS & bearer token file config is used to connect to the actual scrape
  # endpoints for cluster components. This is separate to discovery auth
  # configuration because discovery & scraping are two separate concerns in
  # Prometheus. The discovery auth config is automatic if Prometheus runs inside
  # the cluster. Otherwise, more config options have to be provided within the
  # <kubernetes_sd_config>.
  tls_config:
    ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
  bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token

  kubernetes_sd_configs:
  - role: node

  relabel_configs:
  - action: labelmap
    regex: __meta_kubernetes_node_label_(.+)
  - target_label: __address__
    replacement: kubernetes.default.svc:443
  - source_labels: [__meta_kubernetes_node_name]
    regex: (.+)
    target_label: __metrics_path__
    replacement: /api/v1/nodes/${1}/proxy/metrics

# Scrape config for Kubelet cAdvisor.
#
# This is required for Kubernetes 1.7.3 and later, where cAdvisor metrics
# (those whose names begin with 'container_') have been removed from the
# Kubelet metrics endpoint.  This job scrapes the cAdvisor endpoint to
# retrieve those metrics.
#
# In Kubernetes 1.7.0-1.7.2, these metrics are only exposed on the cAdvisor
# HTTP endpoint; use "replacement: /api/v1/nodes/${1}:4194/proxy/metrics"
# in that case (and ensure cAdvisor's HTTP server hasn't been disabled with
# the --cadvisor-port=0 Kubelet flag).
#
# This job is not necessary and should be removed in Kubernetes 1.6 and
# earlier versions, or it will cause the metrics to be scraped twice.
- job_name: 'kubernetes-cadvisor'

  # Default to scraping over https. If required, just disable this or change to
  # ` + "`" + `http` + "`" + `.
  scheme: https

  # This TLS & bearer token file config is used to connect to the actual scrape
  # endpoints for cluster components. This is separate to discovery auth
  # configuration because discovery & scraping are two separate concerns in
  # Prometheus. The discovery auth config is automatic if Prometheus runs inside
  # the cluster. Otherwise, more config options have to be provided within the
  # <kubernetes_sd_config>.
  tls_config:
    ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
  bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token

  kubernetes_sd_configs:
  - role: node

  relabel_configs:
  - action: labelmap
    regex: __meta_kubernetes_node_label_(.+)
  - target_label: __address__
    replacement: kubernetes.default.svc:443
  - source_labels: [__meta_kubernetes_node_name]
    regex: (.+)
    target_label: __metrics_path__
    replacement: /api/v1/nodes/${1}/proxy/metrics/cadvisor

# Example scrape config for service endpoints.
#
# The relabeling allows the actual service scrape endpoint to be configured
# for all or only some endpoints.
- job_name: 'kubernetes-service-endpoints'

  kubernetes_sd_configs:
  - role: endpoints

  relabel_configs:
  # Example relabel to scrape only endpoints that have
  # "example.io/should_be_scraped = true" annotation.
  #  - source_labels: [__meta_kubernetes_service_annotation_example_io_should_be_scraped]
  #    action: keep
  #    regex: true
  #
  # Example relabel to customize metric path based on endpoints
  # "example.io/metric_path = <metric path>" annotation.
  #  - source_labels: [__meta_kubernetes_service_annotation_example_io_metric_path]
  #    action: replace
  #    target_label: __metrics_path__
  #    regex: (.+)
  #
  # Example relabel to scrape only single, desired port for the service based
  # on endpoints "example.io/scrape_port = <port>" annotation.
  #  - source_labels: [__address__, __meta_kubernetes_service_annotation_example_io_scrape_port]
  #    action: replace
  #    regex: ([^:]+)(?::\d+)?;(\d+)
  #    replacement: $1:$2
  #    target_label: __address__
  #
  # Example relabel to configure scrape scheme for all service scrape targets
  # based on endpoints "example.io/scrape_scheme = <scheme>" annotation.
  #  - source_labels: [__meta_kubernetes_service_annotation_example_io_scrape_scheme]
  #    action: replace
  #    target_label: __scheme__
  #    regex: (https?)
  - action: labelmap
    regex: __meta_kubernetes_service_label_(.+)
  - source_labels: [__meta_kubernetes_namespace]
    action: replace
    target_label: kubernetes_namespace
  - source_labels: [__meta_kubernetes_service_name]
    action: replace
    target_label: kubernetes_name

# Example scrape config for probing services via the Blackbox Exporter.
#
# The relabeling allows the actual service scrape endpoint to be configured
# for all or only some services.
- job_name: 'kubernetes-services'

  metrics_path: /probe
  params:
    module: [http_2xx]

  kubernetes_sd_configs:
  - role: service

  relabel_configs:
  # Example relabel to probe only some services that have "example.io/should_be_probed = true" annotation
  #  - source_labels: [__meta_kubernetes_service_annotation_example_io_should_be_probed]
  #    action: keep
  #    regex: true
  - source_labels: [__address__]
    target_label: __param_target
  - target_label: __address__
    replacement: blackbox-exporter.example.com:9115
  - source_labels: [__param_target]
    target_label: instance
  - action: labelmap
    regex: __meta_kubernetes_service_label_(.+)
  - source_labels: [__meta_kubernetes_namespace]
    target_label: kubernetes_namespace
  - source_labels: [__meta_kubernetes_service_name]
    target_label: kubernetes_name

# Example scrape config for probing ingresses via the Blackbox Exporter.
#
# The relabeling allows the actual ingress scrape endpoint to be configured
# for all or only some services.
- job_name: 'kubernetes-ingresses'

  metrics_path: /probe
  params:
    module: [http_2xx]

  kubernetes_sd_configs:
  - role: ingress

  relabel_configs:
  # Example relabel to probe only some ingresses that have "example.io/should_be_probed = true" annotation
  #  - source_labels: [__meta_kubernetes_ingress_annotation_example_io_should_be_probed]
  #    action: keep
  #    regex: true
  - source_labels: [__meta_kubernetes_ingress_scheme,__address__,__meta_kubernetes_ingress_path]
    regex: (.+);(.+);(.+)
    replacement: ${1}://${2}${3}
    target_label: __param_target
  - target_label: __address__
    replacement: blackbox-exporter.example.com:9115
  - source_labels: [__param_target]
    target_label: instance
  - action: labelmap
    regex: __meta_kubernetes_ingress_label_(.+)
  - source_labels: [__meta_kubernetes_namespace]
    target_label: kubernetes_namespace
  - source_labels: [__meta_kubernetes_ingress_name]
    target_label: kubernetes_name

# Example scrape config for pods
#
# The relabeling allows the actual pod scrape to be configured
# for all the declared ports (or port-free target if none is declared)
# or only some ports.
- job_name: 'kubernetes-pods'

  kubernetes_sd_configs:
  - role: pod

  relabel_configs:
  # Example relabel to scrape only pods that have
  # "example.io/should_be_scraped = true" annotation.
  #  - source_labels: [__meta_kubernetes_pod_annotation_example_io_should_be_scraped]
  #    action: keep
  #    regex: true
  #
  # Example relabel to customize metric path based on pod
  # "example.io/metric_path = <metric path>" annotation.
  #  - source_labels: [__meta_kubernetes_pod_annotation_example_io_metric_path]
  #    action: replace
  #    target_label: __metrics_path__
  #    regex: (.+)
  #
  # Example relabel to scrape only single, desired port for the pod
  # based on pod "example.io/scrape_port = <port>" annotation.
  # Note that __address__ is modified here, so if pod containers' ports
  # are declared, they all will be ignored.
  #  - source_labels: [__address__, __meta_kubernetes_pod_annotation_example_io_scrape_port]
  #    action: replace
  #    regex: ([^:]+)(?::\d+)?;(\d+)
  #    replacement: $1:$2
  #    target_label: __address__
  - action: labelmap
    regex: __meta_kubernetes_pod_label_(.+)
  - source_labels: [__meta_kubernetes_namespace]
    action: replace
    target_label: kubernetes_namespace
  - source_labels: [__meta_kubernetes_pod_name]
    action: replace
    target_label: kubernetes_pod_name
`

func InputIstioPrometheusScrapeConfigs() []*config.ScrapeConfig {
	var scrapeConfigs []*config.ScrapeConfig
	if err := yaml.Unmarshal([]byte(istioScrapeConfigsYaml), &scrapeConfigs); err != nil {
		panic(err)
	}
	return scrapeConfigs
}

// imported from default istio install
var istioScrapeConfigsYaml = `
- job_name: 'istio-mesh'
  # Override the global default and scrape targets from this job every 5 seconds.
  scrape_interval: 5s

  kubernetes_sd_configs:
  - role: endpoints
    namespaces:
      names:
      - istio-system

  relabel_configs:
  - source_labels: [__meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
    action: keep
    regex: istio-telemetry;prometheus


# Scrape config for envoy stats
- job_name: 'envoy-stats'
  metrics_path: /stats/prometheus
  kubernetes_sd_configs:
  - role: pod

  relabel_configs:
  - source_labels: [__meta_kubernetes_pod_container_port_name]
    action: keep
    regex: '.*-envoy-prom'
  - source_labels: [__address__, __meta_kubernetes_pod_annotation_prometheus_io_port]
    action: replace
    regex: ([^:]+)(?::\d+)?;(\d+)
    replacement: $1:15090
    target_label: __address__
  - action: labelmap
    regex: __meta_kubernetes_pod_label_(.+)
  - source_labels: [__meta_kubernetes_namespace]
    action: replace
    target_label: namespace
  - source_labels: [__meta_kubernetes_pod_name]
    action: replace
    target_label: pod_name

  metric_relabel_configs:
  # Exclude some of the envoy metrics that have massive cardinality
  # This list may need to be pruned further moving forward, as informed
  # by performance and scalability testing.
  - source_labels: [ cluster_name ]
    regex: '(outbound|inbound|prometheus_stats).*'
    action: drop
  - source_labels: [ tcp_prefix ]
    regex: '(outbound|inbound|prometheus_stats).*'
    action: drop
  - source_labels: [ listener_address ]
    regex: '(.+)'
    action: drop
  - source_labels: [ http_conn_manager_listener_prefix ]
    regex: '(.+)'
    action: drop
  - source_labels: [ http_conn_manager_prefix ]
    regex: '(.+)'
    action: drop
  - source_labels: [ __name__ ]
    regex: 'envoy_tls.*'
    action: drop
  - source_labels: [ __name__ ]
    regex: 'envoy_tcp_downstream.*'
    action: drop
  - source_labels: [ __name__ ]
    regex: 'envoy_http_(stats|admin).*'
    action: drop
  - source_labels: [ __name__ ]
    regex: 'envoy_cluster_(lb|retry|bind|internal|max|original).*'
    action: drop


- job_name: 'istio-policy'
  # Override the global default and scrape targets from this job every 5 seconds.
  scrape_interval: 5s
  # metrics_path defaults to '/metrics'
  # scheme defaults to 'http'.

  kubernetes_sd_configs:
  - role: endpoints
    namespaces:
      names:
      - istio-system


  relabel_configs:
  - source_labels: [__meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
    action: keep
    regex: istio-policy;http-monitoring

- job_name: 'istio-telemetry'
  # Override the global default and scrape targets from this job every 5 seconds.
  scrape_interval: 5s
  # metrics_path defaults to '/metrics'
  # scheme defaults to 'http'.

  kubernetes_sd_configs:
  - role: endpoints
    namespaces:
      names:
      - istio-system

  relabel_configs:
  - source_labels: [__meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
    action: keep
    regex: istio-telemetry;http-monitoring

- job_name: 'pilot'
  # Override the global default and scrape targets from this job every 5 seconds.
  scrape_interval: 5s
  # metrics_path defaults to '/metrics'
  # scheme defaults to 'http'.

  kubernetes_sd_configs:
  - role: endpoints
    namespaces:
      names:
      - istio-system

  relabel_configs:
  - source_labels: [__meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
    action: keep
    regex: istio-pilot;http-monitoring

- job_name: 'galley'
  # Override the global default and scrape targets from this job every 5 seconds.
  scrape_interval: 5s
  # metrics_path defaults to '/metrics'
  # scheme defaults to 'http'.

  kubernetes_sd_configs:
  - role: endpoints
    namespaces:
      names:
      - istio-system

  relabel_configs:
  - source_labels: [__meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
    action: keep
    regex: istio-galley;http-monitoring
`
