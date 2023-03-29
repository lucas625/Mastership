#!/bin/bash
. ./env_vars.sh || exit 1

export MSC_ROOT_FOLDER="${PWD%/*}" 
export MSC_KUBERNETES_FOLDER="$MSC_ROOT_FOLDER/kubernetes"

cat $MSC_KUBERNETES_FOLDER/istio/peer_authentication.yaml | sed \
    -e "s/\$\$MSC_NAMESPACE/$MSC_NAMESPACE/" |
    kubectl apply -n $MSC_NAMESPACE -f -
cat $MSC_KUBERNETES_FOLDER/istio/deny_other_namespaces.yaml | sed \
    -e "s/\$\$MSC_NAMESPACE/$MSC_NAMESPACE/" |
    kubectl apply -n $MSC_NAMESPACE -f -
cat $MSC_KUBERNETES_FOLDER/istio/allow_nothing.yaml | sed \
    -e "s/\$\$MSC_NAMESPACE/$MSC_NAMESPACE/" |
    kubectl apply -n $MSC_NAMESPACE -f -