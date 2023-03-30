# Hub to pull from. Image will be `Hub/Image:Tag-Variant`
#hub: "gcr.io/istio-testing"
hub: "docker.io/istio"
# Tag to pull from. Image will be `Hub/Image:Tag-Variant`
#TODO tag: "latest"
tag: "1.18.0-alpha.0"

# Variant to pull. Options are "debug" or "distroless". Unset will use the default for the given version.
variant: ""

# Image name to pull from. Image will be `Hub/Image:Tag-Variant`
# If Image contains a "/", it will replace the entire `image` in the pod.
image: ztunnel
#OLD image: proxyv2

# Labels to apply to all top level resources
labels: {}
# Annotations to apply to all top level resources
annotations: {}

# Annotations added to each pod. The default annotations are required for scraping prometheus (in most environments).
podAnnotations:
  prometheus.io/port: "15020"
  prometheus.io/scrape: "true"

# Additional labels to apply on the pod level
podLabels: {}

# Pod resource configuration
resources:
  requests:
    cpu: 500m
    memory: 2048Mi

# List of secret names to add to the service account as image pull secrets
imagePullSecrets: []

# A `key: value` mapping of environment variables to add to the pod
env: {}

# Override for the pod imagePullPolicy
imagePullPolicy: ""

# Settings for multicluster
multiCluster:
  # The name of the cluster we are installing in. Note this is a user-defined name, which must be consistent
  # with Istiod configuration.
  clusterName: ""

# meshConfig defines runtime configuration of components.
# For ztunnel, only defaultConfig is used, but this is nested under `meshConfig` for consistency with other
# components.
# TODO: https://github.com/istio/istio/issues/43248
meshConfig:
  defaultConfig:
    proxyMetadata: {}
