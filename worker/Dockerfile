FROM golang:1.14-alpine as builder
ARG TRADER_HOST
ARG TRADER_PORT
ENV TRADER_HOST=${TRADER_HOST}
ENV TRADER_PORT=${TRADER_PORT}
WORKDIR /go/src/build
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o worker ./cmd

FROM scratch
COPY --from=builder /go/src/build/worker /app/
WORKDIR /app
#EXPOSE $WEB_LISTEN_PORT $RPC_LISTEN_PORT
CMD ["./worker"]