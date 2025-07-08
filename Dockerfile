FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o server .

FROM alpine:3.20

RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime \
    && echo "America/Sao_Paulo" > /etc/timezone

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/.env .
COPY --from=builder /app/log/ ./log/

RUN mkdir -p images log \
    && touch log/log.txt log/last_position.txt

EXPOSE 5000

CMD ["./server"]
