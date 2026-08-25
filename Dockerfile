FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o notifier .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/notifier .
COPY .env .
EXPOSE 8080
CMD ["./notifier"]
