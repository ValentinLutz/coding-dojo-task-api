FROM golang:1.20.2-alpine3.17 AS builder
WORKDIR /build
COPY ./app-gearbox/go.mod ./app-gearbox/go.sum ./
RUN go mod download
COPY ./app-gearbox ./
RUN go build -o app

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /build/app ./
CMD ["./app"]