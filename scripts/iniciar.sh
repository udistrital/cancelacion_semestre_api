#!/bin/bash

set -e -u

echo Iniciando $(date)

tee /tmp/packer.env << EOF
export SUDO=${SUDO}
EOF

if [ -n "${PROXY:-}" ]
then
  tee -a /tmp/packer.env << EOF
export http_proxy=${PROXY}
export HTTP_PROXY=${PROXY}
export https_proxy=${PROXY}
export HTTPS_PROXY=${PROXY}
EOF
fi
