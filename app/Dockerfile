# syntax=docker/dockerfile:1

FROM golang:1.18.4-alpine3.16 AS builder
WORKDIR /build
COPY ./ ./
RUN go build -o app

FROM alpine:3.16.1
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/app ./
CMD ["./app"]
