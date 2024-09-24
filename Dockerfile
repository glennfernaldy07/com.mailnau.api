# Choose whatever you want, version >= 1.16
FROM golang:1.23-alpine as dev

WORKDIR /app

COPY ./ /app/

RUN go install github.com/air-verse/air@latest && \
    go install github.com/pressly/goose/v3/cmd/goose@latest && \
    go mod download

CMD ["air", "-c", ".air.toml"]

