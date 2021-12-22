# Build stage
FROM golang:1.17.5-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o cmd/main cmd/main.go

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/cmd/main .

EXPOSE 8080
CMD [ "/app/cmd/main" ]