FROM golang:1.19.3
WORKDIR /app
COPY . .
RUN go build -o bin/server cmd/main.go
CMD ["./bin/server"]