#!/bin/bash
. ./env_vars.sh || exit 1

export MSC_ROOT_FOLDER="${PWD%/*}" 
export MSC_KUBERNETES_FOLDER="$MSC_ROOT_FOLDER/kubernetes"

# Allow Istio
kubectl label namespace $MSC_NAMESPACE istio-injection=enabled

cat $MSC_KUBERNETES_FOLDER/istio/peer_authentication.yaml | sed \
    -e "s/\$\$MSC_NAMESPACE/$MSC_NAMESPACE/" |
    kubectl apply -n $MSC_NAMESPACE -f -
cat $MSC_KUBERNETES_FOLDER/istio/deny_other_namespaces.yaml | sed \
    -e "s/\$\$MSC_NAMESPACE/$MSC_NAMESPACE/" |
    kubectl apply -n $MSC_NAMESPACE -f -
cat $MSC_KUBERNETES_FOLDER/istio/allow_only_experimenter_for_calculator.yaml | sed \
    -e "s/\$\$MSC_NAMESPACE/$MSC_NAMESPACE/" |
    kubectl apply -n $MSC_NAMESPACE -f -