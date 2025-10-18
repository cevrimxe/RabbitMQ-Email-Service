FROM golang:1.23.3-alpine AS builder

WORKDIR /app

# Publisher kodunu kopyala
COPY publisher/ ./

# Go modüllerini indir
RUN go mod download

# Binary oluştur
RUN CGO_ENABLED=0 GOOS=linux go build -o publisher ./main.go

# Runtime image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/publisher .

EXPOSE 8080

CMD ["./publisher"]
