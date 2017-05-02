#!/bin/bash

set -e -u

packer validate -var sudo= -only docker packer.json
packer build -var sudo= -only docker packer.json

#docker run -ti -v $PWD/..:/root/workspace --name cancelacion_semestre oas0/cancelacion_semestre bash
docker run -ti -v $PWD/..:/root/workspace --name cancelacion_semestre oas0/cancelacion_semestre
