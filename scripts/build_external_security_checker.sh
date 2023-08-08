#!/bin/bash
. ./env_vars.sh || exit 1

export MSC_IMAGE_PULL_POLICY="IfNotPresent"
export MSC_ROOT_FOLDER="${PWD%/*}" 
export MSC_KUBERNETES_FOLDER="$MSC_ROOT_FOLDER/kubernetes"

export MSC_EXTERNAL_NAMESPACE="msc-external"

kubectl delete namespace $MSC_EXTERNAL_NAMESPACE
kubectl create namespace $MSC_EXTERNAL_NAMESPACE

# Un comment to use Istio
# kubectl label namespace $MSC_EXTERNAL_NAMESPACE istio-injection=enabled

cat $MSC_KUBERNETES_FOLDER/security_checker/security_checker.yaml | sed \
    -e "s/\$\$MSC_TARGET_NAMESPACE/$MSC_NAMESPACE/" \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" \
    -e "s/\$\$MSC_SECURITY_CHECKER_SECRET_KEY/$MSC_SECURITY_CHECKER_SECRET_KEY/" | \
    kubectl apply -n $MSC_EXTERNAL_NAMESPACE -f -
