#!/bin/bash

export MSC_LINKERD_FOLDER="$MSC_KUBERNETES_FOLDER/linkerd"

# Set config values
kubectl annotate namespace $MSC_NAMESPACE config.linkerd.io/default-inbound-policy=cluster-authenticated

# Allow Build linkerd elements
cat $MSC_LINKERD_FOLDER/calculator_server.yaml | sed \
    -e "s/\$\$MSC_NAMESPACE/$MSC_NAMESPACE/" |
    kubectl apply -n $MSC_NAMESPACE -f -
cat $MSC_LINKERD_FOLDER/calculator_authorization_policy.yaml | sed \
    -e "s/\$\$MSC_NAMESPACE/$MSC_NAMESPACE/" |
    kubectl apply -n $MSC_NAMESPACE -f -

# Run linkerd
kubectl get -n $MSC_NAMESPACE deploy -o yaml \
  | linkerd inject - \
  | kubectl apply -f -