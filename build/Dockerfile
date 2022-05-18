FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go get github.com/cosmtrek/air@latest
RUN go get github.com/pressly/goose/v3/cmd/goose@latest

COPY . /app

CMD ["air", "-c", ".air.toml"]

