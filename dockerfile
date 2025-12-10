# ---------- STAGE 1: BUILDER ----------
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd/main.go

# ---------- STAGE 2: RUNNER ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/auth-service .

EXPOSE 9000

CMD ["./auth-service"]