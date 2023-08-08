#!/bin/bash

. ./env_vars.sh || exit 1

export MSC_ROOT_FOLDER="${PWD%/*}" 

docker build -t $MSC_TAG_PREFIX/msc-analyzer:$MSC_TAG_VERSION -f $MSC_ROOT_FOLDER/analyzer/Dockerfile $MSC_ROOT_FOLDER/analyzer
docker build -t $MSC_TAG_PREFIX/msc-calculator:$MSC_TAG_VERSION -f $MSC_ROOT_FOLDER/calculator/Dockerfile $MSC_ROOT_FOLDER/calculator
docker build -t $MSC_TAG_PREFIX/msc-experimenter:$MSC_TAG_VERSION -f $MSC_ROOT_FOLDER/experimenter/Dockerfile $MSC_ROOT_FOLDER/experimenter
docker build -t $MSC_TAG_PREFIX/msc-frontend:$MSC_TAG_VERSION -f $MSC_ROOT_FOLDER/frontend/Dockerfile $MSC_ROOT_FOLDER/frontend
docker build -t $MSC_TAG_PREFIX/msc-reverse-proxy:$MSC_TAG_VERSION -f $MSC_ROOT_FOLDER/reverse_proxy/Dockerfile $MSC_ROOT_FOLDER/reverse_proxy
docker build -t $MSC_TAG_PREFIX/msc-security-checker:$MSC_TAG_VERSION -f $MSC_ROOT_FOLDER/security_checker/Dockerfile $MSC_ROOT_FOLDER/security_checker
