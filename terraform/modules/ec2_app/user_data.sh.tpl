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
  reverse_proxy webapp:${port}
}
EOF

# Create Docker Compose file
cat <<EOF > /home/ec2-user/compose.yaml

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

  webapp:
    image: ${app_ecr_repo}:${app_version}
    container_name: webapp
    ports:
      - "${port}:${port}"
    environment:
      - APP_PORT=${port} 
      - APP_AWS_REGION=${aws_region} 
      - APP_AWS_PROFILE=""
      - APP_ENVIRONMENT=${environment}
      - APP_VERSION=${app_version}
    networks:
      - app_network      

networks:
  app_network:
    driver: bridge

volumes:
  caddy_data:
  caddy_config:
EOF

aws ecr get-login-password --region ${aws_region} | docker login --username AWS --password-stdin ${aws_account_id}.dkr.ecr.${aws_region}.amazonaws.com

docker pull ${caddy_ecr_image}
docker pull ${app_ecr_repo}:${app_version}

su ec2-user

docker compose -f /home/ec2-user/compose.yaml up -d