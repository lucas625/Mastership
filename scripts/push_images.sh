#!/bin/bash

. ./env_vars.sh || exit 1

export MSC_ROOT_FOLDER="${PWD%/*}" 

gcloud docker -- push $MSC_TAG_PREFIX/msc-analyzer:$MSC_TAG_VERSION
gcloud docker -- push $MSC_TAG_PREFIX/msc-calculator:$MSC_TAG_VERSION
gcloud docker -- push $MSC_TAG_PREFIX/msc-experimenter:$MSC_TAG_VERSION
gcloud docker -- push $MSC_TAG_PREFIX/msc-frontend:$MSC_TAG_VERSION
gcloud docker -- push $MSC_TAG_PREFIX/msc-reverse-proxy:$MSC_TAG_VERSION
gcloud docker -- push $MSC_TAG_PREFIX/msc-security-checker:$MSC_TAG_VERSION
