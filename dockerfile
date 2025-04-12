FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY config/ ./config/

EXPOSE 8080

CMD ["./main"]
