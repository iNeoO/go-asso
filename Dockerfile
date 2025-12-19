FROM golang:1.22-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /bin/api ./cmd/api

FROM alpine:3.20

WORKDIR /app

RUN addgroup -S app && adduser -S app -G app

COPY --from=builder /bin/api /app/api
COPY --from=builder /src/.env /app/.env

USER app

EXPOSE 5000

ENTRYPOINT ["/app/api"]
