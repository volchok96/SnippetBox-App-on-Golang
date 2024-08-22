FROM golang:1.22.2-alpine AS builder

LABEL stage=intermediate
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/web

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 4000
RUN chmod +x ./main
CMD ["./main"]   