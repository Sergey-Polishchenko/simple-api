FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /server ./cmd/server

FROM alpine:latest

WORKDIR /

COPY --from=builder /server /server
COPY --from=builder /app/.env /.env

CMD ["/server"]
