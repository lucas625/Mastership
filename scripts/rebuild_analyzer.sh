#!/bin/bash
. ./env_vars.sh || exit 1

export MSC_IMAGE_PULL_POLICY="IfNotPresent"
export MSC_ROOT_FOLDER="${PWD%/*}" 
export MSC_KUBERNETES_FOLDER="$MSC_ROOT_FOLDER/kubernetes"

cat $MSC_KUBERNETES_FOLDER/analyzer/analyzer.yaml | sed \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" \
    -e "s/\$\$MSC_ANALYZER_SECRET_KEY/$MSC_ANALYZER_SECRET_KEY/" | \
    kubectl delete -n $MSC_NAMESPACE -f -

cat $MSC_KUBERNETES_FOLDER/analyzer/analyzer.yaml | sed \
    -e "s/\$\$MSC_TAG_PREFIX/$MSC_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$MSC_TAG_VERSION/$MSC_TAG_VERSION/" \
    -e "s/\$\$MSC_IMAGE_PULL_POLICY/$MSC_IMAGE_PULL_POLICY/" \
    -e "s/\$\$MSC_ANALYZER_SECRET_KEY/$MSC_ANALYZER_SECRET_KEY/" | \
    kubectl apply -n $MSC_NAMESPACE -f -
