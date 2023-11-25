FROM golang:1.21.3-bookworm as builder

WORKDIR /app

ADD . /app

RUN go build -o bin/quote .

FROM debian:bookworm-slim

COPY --from=builder /app/bin/quote .

COPY ./migrations /app/migrations

ENV MIGRATIONS_PATH="/app/migrations"

EXPOSE 8200

CMD ["./quote"]