#!/bin/bash

tickers=(PYPL PTON COUP AMD V UBER AAPL PLTR MSFT ZM FB NVDA TWTR SEDG NVTA SPCE TDOC BABA FRHC SQ)
for t in "${tickers[@]}"
do
  docker rm -f rsibot_$t
done