# Start from the official Golang image
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o api-server main.go

# Use a minimal image for running
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api-server .
EXPOSE 8000
ENTRYPOINT ["./api-server"]
