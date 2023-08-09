FROM golang:1.21-alpine3.18 AS builder
WORKDIR /build
COPY ./app/go.mod ./app/go.sum ./
RUN go mod download
COPY ./app ./
RUN go build -o app

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /build/app ./
CMD ["./app"]