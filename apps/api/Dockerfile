FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./cmd

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates netcat-openbsd
WORKDIR /root/

COPY --from=builder /app/api .

EXPOSE 3000

CMD ["./api"]