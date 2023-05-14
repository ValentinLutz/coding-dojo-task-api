FROM golang:1.20.4-alpine3.17 AS builder
WORKDIR /build
COPY ./app-fiber/go.mod ./app-fiber/go.sum ./
RUN go mod download
COPY ./app-fiber ./
RUN go build -o app

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /build/app ./
CMD ["./app"]