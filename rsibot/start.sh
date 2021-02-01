#!/bin/bash

./stop.sh

source ./rsibot.env
docker build -t rsibot .

tickers=(PYPL PTON COUP AMD V UBER AAPL PLTR MSFT ZM FB NVDA TWTR SEDG NVTA SPCE TDOC BABA FRHC SQ)
for t in "${tickers[@]}"
do
  source ./tickers/$t.env &&
  docker run --name rsibot_$TICKER -e TICKER=$TICKER -e RSI_OVERSOLD=$RSI_OVERSOLD -e TELEGRAM_BOT_TOKEN=$TELEGRAM_BOT_TOKEN -e TELEGRAM_CHANNEL_NAME=$TELEGRAM_CHANNEL_NAME -e TINKOFF_TOKEN=$TINKOFF_TOKEN -d rsibot
done