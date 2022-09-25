# Base image
FROM golang:1.18-alpine

RUN apk add --no-cache git

WORKDIR /api

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

# Build the Go API
RUN go build -o main .

EXPOSE 5000

CMD ["/api/main"]