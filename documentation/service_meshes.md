# Service Meshes

This document offers guides on how to setup and use the service meshes Linkerd and Istio for the Lucas Aurelio's msc.

## Table of Contents

- [Service Meshes](#service-meshes)
  - [Table of Contents](#table-of-contents)
  - [Installations](#installations)
    - [Linkerd Installation](#linkerd-installation)
      - [Install Linkerd CLI](#install-linkerd-cli)
      - [Install Linkerd on Cluster](#install-linkerd-on-cluster)
    - [Istio Installation](#istio-installation)
  - [Testing](#testing)
    - [Demo App](#demo-app)
    - [Mesh The Application](#mesh-the-application)
      - [Mesh With Linkerd](#mesh-with-linkerd)
    - [Mesh With Istio](#mesh-with-istio)
  - [Extra](#extra)
    - [Istio](#istio)
  - [Test TLS](#test-tls)
    - [Test Istio TLS](#test-istio-tls)
    - [Test Linkerd TLS](#test-linkerd-tls)

## Installations

### Linkerd Installation

#### Install Linkerd CLI

```bash
# Install the CLI (It might be necessary to add to PATH afterwards)
curl --proto '=https' --tlsv1.2 -sSfL https://run.linkerd.io/install | sh

# Check version
linkerd version
```

#### Install Linkerd on Cluster

```bash
# Verifies if cluster can install Linkerd
linkerd check --pre

# Install Linkerd crds
linkerd install --crds | kubectl apply -f -

# Apply linkerd to cluster
linkerd install --set proxyInit.runAsRoot=true | kubectl apply -f -

# Verify installation on cluster
linkerd check
```

### Istio Installation

```bash
# Download Istio
curl -L https://istio.io/downloadIstio | sh -

# Go to the folder (in this case version 1.15.1)
cd istio-1.15.1

# Add to PATH
export PATH=$PWD/bin:$PATH

# Install Istio using the demo profile
istioctl install --set profile=demo -y

# It might be necessary to install like this instead
istioctl install --set values.global.proxy.privileged=true
```

## Testing

### Demo App

It is important to be able to test the installation of the service meshes. As such, install the demo app and check it running.

```bash
# Install the demo app
curl --proto '=https' --tlsv1.2 -sSfL https://run.linkerd.io/emojivoto.yml | kubectl apply -f -

# Forward to port
kubectl -n emojivoto port-forward svc/web-svc 8080:80
```

### Mesh The Application

The topics below show how to add the demo app to the specific meshes. Remember to delete the pods if you apply to one mesh and then want to apply to another.

#### Mesh With Linkerd

```bash
# Adds the proxy container to the pods
kubectl get -n emojivoto deploy -o yaml \
  | linkerd inject - \
  | kubectl apply -f -

# Check the data plane
linkerd -n emojivoto check --proxy

# Install viz for cluster metrics
linkerd viz install | kubectl apply -f -

# Access dashboard
linkerd viz dashboard &
```

### Mesh With Istio

```bash
# Create namespace
kubectl create namespace emojivoto

# Enable Istio for namespace 
kubectl label namespace emojivoto istio-injection=enabled
# kubectl label namespace emojivoto istio-injection=disabled --overwrite To disable

# You will be required to update the pods or recreate them
```

## Extra

### Istio

```bash
# Display authorization policies
kubectl get authorizationpolicy.security.istio.io -n mastership
```

## Test TLS

### Test Istio TLS

```bash
# Install prometheus for istio
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.18/samples/addons/prometheus.yaml

# Open the dashboards
istioctl dashboard prometheus

# Example of Prometheus query:
istio_requests_total{namespace="mastership"}[1h]
```

### Test Linkerd TLS

```bash
# Check if everything is protected by TLS
linkerd viz -n mastership edges deployment

# Checks the traffic coming in and out of the deployments
linkerd viz -n mastership tap deployment
```
