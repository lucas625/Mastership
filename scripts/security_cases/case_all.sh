#!/bin/bash

export MSC_MESH=$1
export MSC_ISTIO_MESH="istio"
export MSC_LINKERD_MESH="linkerd"
export MSC_NO_MESH="no-mesh"

if [ $MSC_MESH = $MSC_ISTIO_MESH ] || [ $MSC_MESH = $MSC_LINKERD_MESH] || [ $MSC_MESH = $MSC_NO_MESH ]; then
    echo "Running for $MSC_MESH"
else
    echo "Unknown mesh: $MSC_MESH. Exiting"
    exit 1
fi

. ./base.sh

if [ $MSC_MESH = $MSC_ISTIO_MESH ]; then
    echo "Configuring Istio"
    $MSC_SCRIPTS_FOLDER/configure_istio.sh
fi

echo "Creating kubernetes elements"

cat $MSC_KUBERNETES_FOLDER/analyzer/analyzer.yaml | sed \
    -e "s/\$\$MSC_NAMESPACE/$MSC_NAMESPACE/" \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" \
    -e "s/\$\$MSC_ANALYZER_SECRET_KEY/$MSC_ANALYZER_SECRET_KEY/" | \
    kubectl apply -n $MSC_NAMESPACE -f -

cat $MSC_KUBERNETES_FOLDER/calculator/calculator.yaml | sed \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" | \
    kubectl apply -n $MSC_NAMESPACE -f -

cat $MSC_KUBERNETES_FOLDER/experimenter/experimenter.yaml | sed \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" | \
    kubectl apply -n $MSC_NAMESPACE -f -

cat $MSC_KUBERNETES_FOLDER/frontend/frontend.yaml | sed \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" | \
    kubectl apply -n $MSC_NAMESPACE -f -

cat $MSC_KUBERNETES_FOLDER/reverse_proxy/reverse_proxy.yaml | sed \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" \
    -e "s/\$\$MSC_MINIKUBE_CLUSTER_IP/$MSC_MINIKUBE_CLUSTER_IP/" | \
    kubectl apply -n $MSC_NAMESPACE -f -

cat $MSC_KUBERNETES_FOLDER/security_checker/security_checker.yaml | sed \
    -e "s/\$\$MSC_TARGET_NAMESPACE/$MSC_TARGET_NAMESPACE/" \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" \
    -e "s/\$\$MSC_SECURITY_CHECKER_SECRET_KEY/$MSC_SECURITY_CHECKER_SECRET_KEY/" | \
    kubectl apply -n $MSC_NAMESPACE -f -

if [ $MSC_MESH = $MSC_LINKERD_MESH ]; then
    echo "Configuring Linkerd"
    $MSC_SCRIPTS_FOLDER/configure_linkerd.sh
fi

. ./build_external_security_checker.sh
