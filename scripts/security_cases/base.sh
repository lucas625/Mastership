#!/bin/bash

export MSC_ROOT_FOLDER=$(dirname ${PWD%/*})
export MSC_SCRIPTS_FOLDER="$MSC_ROOT_FOLDER/scripts"
export MSC_KUBERNETES_FOLDER="$MSC_ROOT_FOLDER/kubernetes"
export MSC_IMAGE_PULL_POLICY="IfNotPresent"
export MSC_EXTERNAL_NAMESPACE="msc-external"

. $MSC_SCRIPTS_FOLDER/env_vars.sh || exit 1

export MSC_TARGET_NAMESPACE=$MSC_NAMESPACE

echo "Removing kubernetes elements"
kubectl delete namespace $MSC_NAMESPACE
kubectl delete namespace $MSC_EXTERNAL_NAMESPACE

echo "Creating $MSC_NAMESPACE namespace"
kubectl create namespace $MSC_NAMESPACE

echo "Creating service account"
kubectl create serviceaccount msc-experimenter-service-account -n $MSC_NAMESPACE
