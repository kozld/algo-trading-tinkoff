FROM golang:1.14-alpine as builder
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
ARG RPC_LISTEN_PORT
ENV RPC_LISTEN_PORT=${RPC_LISTEN_PORT}
WORKDIR /go/src/build
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o trader ./cmd

FROM scratch
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=builder /go/src/build/trader /app/
WORKDIR /app
#EXPOSE $WEB_LISTEN_PORT $RPC_LISTEN_PORT
CMD ["./trader"]