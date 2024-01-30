#!/bin/bash


bash $(pwd)/solutions/setup-kind-clusters.sh $1
bash $(pwd)/solutions/ocm-join-clusters.sh $2
bash $(pwd)/solutions/setup-rcs-deployer.sh $1 $2 $3