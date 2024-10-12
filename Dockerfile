FROM golang:1.23.2

WORKDIR /app/common

COPY go.mod go.sum ./
RUN go mod download

COPY . .

