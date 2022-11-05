FROM golang:1.19.3
WORKDIR /app
COPY . .
RUN go build -o bin/server src/cmd/main.go
CMD ["./bin/server"]