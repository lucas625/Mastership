#!/bin/bash

export MSC_ISTIO_FOLDER="$MSC_KUBERNETES_FOLDER/istio"

# Allow Istio
kubectl label namespace $MSC_NAMESPACE istio-injection=enabled

cat $MSC_ISTIO_FOLDER/telemetry.yaml |
    kubectl apply -f -
cat $MSC_ISTIO_FOLDER/peer_authentication.yaml | sed \
    -e "s/\$\$MSC_NAMESPACE/$MSC_NAMESPACE/" |
    kubectl apply -n $MSC_NAMESPACE -f -
cat $MSC_ISTIO_FOLDER/deny_other_namespaces.yaml | sed \
    -e "s/\$\$MSC_NAMESPACE/$MSC_NAMESPACE/" |
    kubectl apply -n $MSC_NAMESPACE -f -
cat $MSC_ISTIO_FOLDER/allow_only_experimenter_for_calculator.yaml | sed \
    -e "s/\$\$MSC_NAMESPACE/$MSC_NAMESPACE/" |
    kubectl apply -n $MSC_NAMESPACE -f -