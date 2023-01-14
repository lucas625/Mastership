#!/bin/bash

. ./env_vars.sh || exit 1

docker build -t $MSC_TAG_PREFIX/msc-analyzer:$MSC_TAG_VERSION -f ../analyzer/Dockerfile analyzer
docker build -t $MSC_TAG_PREFIX/msc-calculator:$MSC_TAG_VERSION -f ../calculator/Dockerfile calculator
docker build -t $MSC_TAG_PREFIX/msc-experimenter:$MSC_TAG_VERSION -f ../experimenter/Dockerfile experimenter
docker build -t $MSC_TAG_PREFIX/msc-frontend:$MSC_TAG_VERSION -f ../frontend/Dockerfile frontend
docker build -t $MSC_TAG_PREFIX/msc-reverse-proxy:$MSC_TAG_VERSION -f ../reverse_proxy/Dockerfile reverse_proxy
