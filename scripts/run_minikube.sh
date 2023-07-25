#!/bin/bash

minikube start || exit 1
minikube config set cpus 6
minikube config set memory 8192
eval $(minikube docker-env) || exit 1

. ./env_vars.sh || exit 1

./build.sh

sleep 10s

kubectl -n $MSC_NAMESPACE port-forward svc/msc-reverse-proxy 8080:80 || exit 1
