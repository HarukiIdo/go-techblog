# syntax=docker/dockerfile:1
FROM golang:1.17

# コンテナの作業ディレクトリを設定
WORKDIR /app

# ローカルのgo.modとgo.sumを作業ディレクトリにコピー
COPY go.mod go.sum ./

# 必要なパッケージをインストール
RUN go mod download

# airをインストール
RUN go install github.com/cosmtrek/air@latest

# カレントディレクトリを作業ディレクトリにコピー
COPY . ./

# ビルド
RUN go build -o app 

# アプリを起動
CMD ["app"]

