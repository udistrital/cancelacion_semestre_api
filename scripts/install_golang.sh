#!/bin/bash
echo 'Ejecutando: install_golang.sh'

set -e -u

[ -e /tmp/packer.env ] && source /tmp/packer.env

if go version&>/dev/null
then
  echo 'Go ya está instalado. Nada que hacer.'
else

$SUDO yum install -y golang
$SUDO tee /etc/profile.d/go.sh << 'EOF'
export GOROOT=/usr/lib/golang
export GOBIN=$GOROOT/bin
export GOPATH=/root
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
EOF

# usergolang=vagrant
# homeuser=/home/$usergolang
# file=$homeuser/.bashrc
# tee $file << 'EOF'
# # Golang Path
# export GOROOT=/usr/lib/golang
# export GOBIN=$GOROOT/bin
# export GOPATH=$HOME
# export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
# EOF
# $SUDO chown $usergolang:$usergolang $file
# source ~/.bashrc

source /etc/profile.d/go.sh

$SUDO ldconfig
go version
go env

fi
