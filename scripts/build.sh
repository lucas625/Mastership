#!/bin/bash

./remove_kubernetes_elements.sh

echo "Building the images"
./build_images.sh || exit 1

echo "Creating kubernetes elements"
./create_kubernetes_elements.sh || exit 1