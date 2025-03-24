# Stage 1: Build the application
FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/main.go

# Stage 2: Create a lightweight image
FROM alpine
WORKDIR /build
COPY --from=builder /build/main /build/main
RUN addgroup -S appuser && adduser -S appuser -G appuser
RUN chown -R appuser:appuser /build
USER appuser
RUN rm -rf /var/cache/apk/*
CMD ["./main"]
