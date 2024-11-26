FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o myapp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp ./

COPY .env .env

COPY config/migrations config/migrations

ENV PORT=$HTTP_SERVER_ADDRESS

EXPOSE $PORT

CMD ["./myapp"]
