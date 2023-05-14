FROM golang:1.20.2-alpine3.17 AS builder
WORKDIR /build
COPY ./app-chi/go.mod ./app-chi/go.sum ./
RUN go mod download
COPY ./app-chi ./
RUN go build -o app

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /build/app ./
CMD ["./app"]