#!/bin/bash
. ./env_vars.sh || exit 1

export MSC_IMAGE_PULL_POLICY="IfNotPresent"
export MSC_ROOT_FOLDER="${PWD%/*}" 
export MSC_KUBERNETES_FOLDER="$MSC_ROOT_FOLDER/kubernetes"
export MSC_EXTERNAL_NAMESPACE="msc-external"

kubectl delete serviceaccount msc-experimenter-service-account -n $MSC_NAMESPACE

kubectl delete namespace $MSC_NAMESPACE
kubectl delete namespace $MSC_EXTERNAL_NAMESPACE
