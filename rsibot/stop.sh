#!/bin/bash

docker ps -a | grep "rsibot_*" | awk '{print $10}' | xargs docker rm -f