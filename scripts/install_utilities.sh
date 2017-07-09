#!/bin/bash
echo 'Ejecutando: install_utilities.sh'

set -e -u

[ -e /tmp/packer.env ] && source /tmp/packer.env

#rationale: list of aplications to install
list=(vim git nmap tree wget)
#rationale: list of packages (no bin app entry)
list_packages=(mlocate)

install=installed
for p in ${list[*]}
do
  if which $p&>/dev/null
  then
    echo "$p ya está instalado."
  else
    install=no_installed
  fi
done

for p in ${list_packages[*]}
do
  if yum list installed $p&>/dev/null
  then
    echo "El paquete $p ya está instalado."
  else
    install=no_installed
  fi
done

if [ "$install" = "installed" ]
then
  echo 'Utilities ya están instalados. Nada que hacer.'
else
  $SUDO yum install -y ${list[*]} ${list_packages[*]}
fi

$SUDO yum install -y epel-release
