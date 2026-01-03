FROM golang:1.25.2-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

FROM alpine:latest

WORKDIR /app

RUN addgroup -g 1001 -S appuser && \
    adduser -u 1001 -S appuser -G appuser

COPY --from=builder --chown=appuser:appuser /app/main .

COPY --chown=appuser:appuser web/ ./web/

USER appuser

EXPOSE 8080

CMD ["./main"]