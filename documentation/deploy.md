# Deploy on gcloud

- [Deploy on gcloud](#deploy-on-gcloud)
  - [Setup the environment variables](#setup-the-environment-variables)
  - [Point to correct project](#point-to-correct-project)
  - [Building the images](#building-the-images)
  - [Setup cluster](#setup-cluster)
  - [Create the static IP](#create-the-static-ip)
  - [Setup pods and services](#setup-pods-and-services)
    - [Service Meshes](#service-meshes)

## Setup the environment variables

```bash
# Build the environment variables.
./build_env_vars.sh

# Run the environment variables
# Remember to change the values as needed.
. ./env_vars.sh
```

## Point to correct project

```bash
# Performing login.
gcloud auth login
gcloud config set project Mastership
```

## Building the images

```bash
# Go to scripts folder
cd scripts

# Build the docker images
./build_images.sh

# Push them to the cloud repository
./push_images.sh
```

## Setup cluster

```bash
# Creates the cluster.
# Remember to go to GKE and delete the cluster when it is no longer necessary.
gcloud container clusters create msc-custom \
  --machine-type custom-4-8192 \
  --zone southamerica-east1-a \
  --num-nodes 1

# Update `kubectl` context in case the current context be not pointing to the cluster created.
kubectl config current-context
kubectl config get-contexts
kubectl config use-context [CONTEXT_NAME]

# Ensure connection to cluster.
gcloud container clusters get-credentials msc-e2-standard-2 --zone southamerica-east1-a

# Resize the cluster to use a single node.
gcloud container clusters resize msc-e2-standard-2 --zone southamerica-east1-a --size=1

# For later deleting the cluster
gcloud container clusters delete msc-custom --zone southamerica-east1-a
```

## Create the static IP

```bash
# Creates the static ip.
gcloud compute addresses create msc-static-ip --region southamerica-east1

# Check the IP.
gcloud compute addresses describe msc-static-ip --region southamerica-east1

# Delete the static IP when not using.
gcloud compute addresses delete msc-static-ip --region southamerica-east1
```

## Setup pods and services

```bash
# Run kubernetes.
./create_kubernetes_elements.sh
```

### Service Meshes

Please check the [service meshes documentation](service_meshes.md).
