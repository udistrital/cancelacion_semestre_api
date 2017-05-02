#!/bin/bash
echo 'Ejecutando: provides_docker.sh'

set -e -u

[ -e /tmp/packer.env ] && source /tmp/packer.env
