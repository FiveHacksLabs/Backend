# Gunakan base image resmi untuk Golang (misalnya versi 1.19)
FROM golang:1.19-alpine AS builder

# Atur working directory
WORKDIR /app

# Copy go.mod dan go.sum, lalu download dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy source code aplikasi
COPY . .

# Build binary
RUN CGO_ENABLED=0 go build -o main .

# Stage kedua: minimal image untuk menjalankan binary
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# Expose port (sesuaikan dengan port yang digunakan aplikasi, misalnya 8080)
EXPOSE 8080

# Jalankan binary
CMD ["./main"]
