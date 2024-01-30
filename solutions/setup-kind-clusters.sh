#!/bin/bash

cd $(dirname ${BASH_SOURCE})

set -e

initialize_kind() {
    if [ -x "$1" ]; then
        kind="$1"
    else
        kind="/usr/local/bin/clusteradm"
    fi
}

initialize_kind "$1"

# Check if kind is executable
if [ ! -x "$kind" ]; then
    echo "Error: The specified kind binary is not executable or does not exist." >&2
    exit 1
fi

hub=${CLUSTER1:-hub}
c1=${CLUSTER1:-cluster1}
c2=${CLUSTER2:-cluster2}

hubctx="kind-${hub}"
c1ctx="kind-${c1}"
c2ctx="kind-${c2}"

"${kind}" create cluster --name "${hub}"
"${kind}" create cluster --name "${c1}"
"${kind}" create cluster --name "${c2}"