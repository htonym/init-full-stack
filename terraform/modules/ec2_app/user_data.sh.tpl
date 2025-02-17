#!/bin/bash

# Assumes docker and docker compose are already installed on AL2023 AMI. If not
# install here.

# Create Caddyfile with dynamic values like domain.
# REMOVE acme_ca for production. Otherwise staging cert is used.
cat <<EOF > /home/ec2-user/Caddyfile

{
  acme_ca https://acme-staging-v02.api.letsencrypt.org/directory
}

${sub_domain} {
  root * /usr/share/caddy
  file_server
}
EOF

# Create Docker Compose file
cat <<EOF > /home/ec2-user/docker-compose.yml

services:
  caddy:
    image: ${caddy_ecr_image}
    container_name: caddy
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile
      - caddy_data:/data
      - caddy_config:/config
    networks:
      - app_network

networks:
  app_network:

volumes:
  caddy_data:
  caddy_config:
EOF

aws ecr get-login-password --region ${aws_region} | docker login --username AWS --password-stdin ${aws_account_id}.dkr.ecr.${aws_region}.amazonaws.com

docker pull ${caddy_ecr_image}

su ec2-user

docker compose -f /home/ec2-user/docker-compose.yml up -d