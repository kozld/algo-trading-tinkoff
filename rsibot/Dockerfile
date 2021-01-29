FROM golang:1.14-alpine as builder
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
WORKDIR /go/src/build
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rsibot ./cmd

FROM scratch
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=builder /go/src/build/rsibot /app/
WORKDIR /app
CMD ["./rsibot"]