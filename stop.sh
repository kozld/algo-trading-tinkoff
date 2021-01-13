#!/bin/bash

docker-compose down
docker rm -f database
docker rm -f trader
docker rm -f worker
docker rm -f grabber