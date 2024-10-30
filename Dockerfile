FROM golang:1.21-alpine

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o main ./cmd/blockchain_test/main.go

EXPOSE 8080
CMD ["./main"]