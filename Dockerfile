FROM golang:1.25.1-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o server cmd/wecare/main.go

from alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8000
CMD ["./server"]