FROM rust:1-alpine3.17 AS builder
RUN apk add --no-cache musl-dev

WORKDIR /build
COPY ./app/Cargo.toml ./app/Cargo.lock ./
COPY ./app/src ./src
RUN cargo build --release --bin app-axum

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /build/target/release/app-axum ./
CMD ["./app-axum"]