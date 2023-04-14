#!/bin/sh

# get latest update from repository
git pull origin master

# build & run application
docker compose -f prod-docker-compose.yaml up --build -d

# clear unused images
docker image prune
