#!/bin/bash

echo "Run and create container (image should be exist)"
# docker run -d --name dfpContainer -p 3000:8080 alta-shop-project-docker:latest

docker run -d --name dfpContainer\
  -p 80:80 \
  -e "HTTP_PORT=:80" \
  -e "MYSQL_ROOT_PASSWORD=12345678" \
  -e "MYSQL_USER=root" \
  -e "MYSQL_PASSWORD=root" \
  -e "MYSQL_DATABASE=alta_shop_project" \
  -e "MYSQL_HOST=13.213.1.38" \
  -e "MYSQL_PORT=3306" \
  fawwazaf/project-alta-deployment:latest
