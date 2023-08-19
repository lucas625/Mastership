#!/bin/bash

export MSC_MESH=$1
export MSC_ISTIO_MESH="istio"
export MSC_LINKERD_MESH="linkerd"

if [ $MSC_MESH = $MSC_ISTIO_MESH ] || [ $MSC_MESH = $MSC_LINKERD_MESH]; then
    echo "Running for $MSC_MESH"
else
    echo "Unknown mesh: $MSC_MESH. Exiting"
    exit 1
fi
