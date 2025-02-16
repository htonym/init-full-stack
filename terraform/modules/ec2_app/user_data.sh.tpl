#!/bin/bash

docker pull caddy

echo "${sub_domain} {
  root * /usr/share/caddy
  file_server
}" > /home/ec2-user/Caddyfile

docker run -d --name caddy \
    -p 80:80 \
    -p 443:443 \
    -v /home/ec2-user/Caddyfile:/etc/caddy/Caddyfile \
    -v caddy_data:/data \
    caddy