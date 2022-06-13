package main

import (
	"log"
	"os"

	"github.com/HarukiIdo/go-techblog/db"
	"github.com/HarukiIdo/go-techblog/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	// 環境変数の読み込み
	if err := godotenv.Load(); err != nil {
		log.Println(".envファイルの読み込みに失敗しました")
	}
}

func main() {

	// Echoインスタンスを作成
	e := echo.New()

	// src/cssを/cssのパスで,
	// src/jsを/jsのパスでアクセス可能にする
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	// DB接続設定
	db := db.NewDB(e)
	defer db.Close()

	// ルーティング
	router.NewRouter(e, db)

	// ポートの指定
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// サーバ起動
	e.Logger.Fatal(e.Start(":" + port))
}
