# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o time-service .

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/time-service .

EXPOSE 9090

ENTRYPOINT ["./time-service"]
