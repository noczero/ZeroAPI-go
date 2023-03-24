#!/bin/sh

# get latest update from repository
git pull origin master

# build & run application
docker compose up --build -d
