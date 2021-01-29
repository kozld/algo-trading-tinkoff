#!/bin/bash

docker rm -f rsibot
source ./rsibot.env

docker build -t rsibot . &&
docker run --name rsibot_$TICKER -e TICKER=$TICKER -e RSI_OVERSOLD=$RSI_OVERSOLD -e TELEGRAM_BOT_TOKEN=$TELEGRAM_BOT_TOKEN -e TELEGRAM_CHANNEL_NAME=$TELEGRAM_CHANNEL_NAME -e TINKOFF_TOKEN=$TINKOFF_TOKEN -d rsibot