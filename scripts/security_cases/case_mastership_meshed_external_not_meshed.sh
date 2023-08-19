#!/bin/bash

. ./validate.sh

. ./base.sh

echo "Creating kubernetes elements"

if [ $MSC_MESH = $MSC_ISTIO_MESH ]; then
    echo "Configuring Istio"
    $MSC_SCRIPTS_FOLDER/configure_istio.sh
fi

cat $MSC_KUBERNETES_FOLDER/calculator/calculator.yaml | sed \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" | \
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

export MSC_MESH="no-mesh"
. ./build_external_security_checker.sh
