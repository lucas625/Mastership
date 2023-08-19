#!/bin/bash

echo "Building external security checker for $MSC_MESH"

export MSC_EXTERNAL_NAMESPACE="msc-external"

kubectl create namespace $MSC_EXTERNAL_NAMESPACE

if [ $MSC_MESH = $MSC_ISTIO_MESH ]; then
    kubectl label namespace $MSC_EXTERNAL_NAMESPACE istio-injection=enabled
fi

cat $MSC_KUBERNETES_FOLDER/security_checker/security_checker.yaml | sed \
    -e "s/\$\$MSC_TARGET_NAMESPACE/$MSC_NAMESPACE/" \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" \
    -e "s/\$\$MSC_SECURITY_CHECKER_SECRET_KEY/$MSC_SECURITY_CHECKER_SECRET_KEY/" | \
    kubectl apply -n $MSC_EXTERNAL_NAMESPACE -f -

if [ $MSC_MESH = $MSC_LINKERD_MESH ]; then
    kubectl get -n $MSC_EXTERNAL_NAMESPACE deploy -o yaml \
      | linkerd inject - \
      | kubectl apply -f -
fi
