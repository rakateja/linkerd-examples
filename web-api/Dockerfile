FROM golang:1.20 AS builder

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o web-api main.go
  
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/web-api .
ENTRYPOINT ./web-api
