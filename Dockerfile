FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . . 

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o api .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/api /app/api
COPY --from=builder /app/api-docs /app/api-docs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/app/api"]