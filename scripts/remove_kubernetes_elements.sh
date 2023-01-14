#!/bin/bash
. ./env_vars.sh || exit 1

export MSC_IMAGE_PULL_POLICY="IfNotPresent"
export MSC_KUBERNETES_FOLDER="../kubernetes"

cat $MSC_KUBERNETES_FOLDER/analyzer/analyzer.yaml | sed \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" \
    -e "s/\$\$MSC_ANALYZER_SECRET_KEY/$MSC_ANALYZER_SECRET_KEY/" | \
    kubectl delete -n $MSC_NAMESPACE -f -

cat $MSC_KUBERNETES_FOLDER/calculator/calculator.yaml | sed \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" | \
    kubectl delete -n $MSC_NAMESPACE -f -

cat $MSC_KUBERNETES_FOLDER/experimenter/experimenter.yaml | sed \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" | \
    kubectl delete -n $MSC_NAMESPACE -f -

cat $MSC_KUBERNETES_FOLDER/frontend/frontend.yaml | sed \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" | \
    kubectl delete -n $MSC_NAMESPACE -f -

cat $MSC_KUBERNETES_FOLDER/reverse_proxy/reverse_proxy_minikube.yaml | sed \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" \
    -e "s/\$\$MSC_MINIKUBE_CLUSTER_IP/$MSC_MINIKUBE_CLUSTER_IP/" | \
    kubectl delete -n $MSC_NAMESPACE -f -

kubectl delete namespace $MSC_NAMESPACE
