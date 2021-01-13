#!/bin/bash

./stop.sh

docker-compose build --no-cache
docker-compose up -d database
sleep 10
docker-compose up -d trader
docker-compose up -d worker
docker-compose run grabber python main.py shell