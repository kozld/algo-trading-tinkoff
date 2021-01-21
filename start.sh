#!/bin/bash

./stop.sh

docker-compose build --no-cache
sleep 10
docker-compose up -d trader
docker-compose up -d worker
docker-compose run -d grabber python grabber.py shell