#!/bin/bash

minikube start || exit 1
minikube config set cpus 6
minikube config set memory 8192
eval $(minikube docker-env) || exit 1

./build_images.sh || exit 1
./create_kubernetes_elements.sh