#TAHAP 1: BUILD
FROM golang:1.24 AS builder

WORKDIR /app

COPY app/go.mod ./
COPY app/go.sum ./
RUN go mod download

COPY app/. ./
RUN go build -o myapp

#TAHAP 2: RUNTIME MINIMAL
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/myapp .

EXPOSE 8080
CMD ["./myapp"]
