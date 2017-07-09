#!/bin/bash
echo 'Ejecutando: install_nginx.sh'

set -e -u

[ -e /tmp/packer.env ] && source /tmp/packer.env

# rationale: instalar nginx
if nginx -v &> /dev/null
then
  echo 'NGINX ya está instalado, nada que hacer.'
else
  $SUDO yum install -y nginx
  $SUDO systemctl enable nginx
fi

# rationale: configurar api y cliente
path=/etc/nginx/projects.d
file=$path/cancelacion_semestre.conf
if [ -f $file ]
then
  echo "El archivo $file ya existe. Nada que hacer."
else
  $SUDO mkdir $path
  $SUDO tee $file << 'EOF'
location /cancelacion_semestre {
    alias /usr/share/nginx/html/cancelacion_semestre_cliente/app;
}

location /cancelacion_semestre/v1 {
    proxy_pass http://localhost:8080;
}
EOF

fi

$SUDO systemctl start httpd
