#!/bin/bash

export MSC_ROOT_FOLDER="${PWD%/*}" 
export MSC_SCRIPTS_FOLDER="$MSC_ROOT_FOLDER/scripts"
export MSC_KUBERNETES_FOLDER="$MSC_ROOT_FOLDER/kubernetes"
export MSC_IMAGE_PULL_POLICY="IfNotPresent"

. $MSC_SCRIPTS_FOLDER/env_vars.sh || exit 1

export MSC_TARGET_NAMESPACE=$MSC_NAMESPACE

echo "Removing kubernetes elements"
$MSC_SCRIPTS_FOLDER/remove_kubernetes_elements.sh

echo "Creating $MSC_NAMESPACE namespace"
kubectl create namespace $MSC_NAMESPACE

echo "Creating service account"
kubectl create serviceaccount msc-experimenter-service-account -n $MSC_NAMESPACE
