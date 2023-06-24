FROM golang:1.20.5-alpine3.17 AS builder
WORKDIR /build
COPY ./app/go.mod ./app/go.sum ./
RUN go mod download
COPY ./app ./
RUN go build -o app

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /build/app ./
CMD ["./app"]