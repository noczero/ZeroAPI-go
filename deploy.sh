#!/bin/sh

# go to the driectory
PROJECT_PATH=$(pwd)
cd PROJECT_PATH

# get latest update from repository
git pull origin master

# build & run application
docker compose up --build -d
