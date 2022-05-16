FROM golang:1.17

COPY go.mod /app/go.mod
COPY go.mod /app/go.sum

WORKDIR /app

COPY . /app

RUN go mod download

