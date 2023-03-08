FROM golang:1.20.2-alpine3.17 AS builder
WORKDIR /build
COPY ./app/go.mod ./app/go.sum ./
RUN go mod download
COPY ./app ./
RUN go build -o app

FROM alpine:3.17
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/app ./
CMD ["./app"]
