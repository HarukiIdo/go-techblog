FROM golang:1.17

# コンテナの作業ディレクトリを設定
WORKDIR /app

# ローカルのgo.modとgo.sumを作業ディレクトリにコピー
COPY go.mod go.sum .

# 必要なパッケージをインストール
# airをインストール
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
COPY . .


# air -c .air.tomlコマンドでairを起動する
CMD ["main"]

