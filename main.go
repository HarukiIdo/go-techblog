package main

import (
	"os"

	"github.com/HarukiIdo/go-techblog/db"
	"github.com/HarukiIdo/go-techblog/handler/router"
	"github.com/HarukiIdo/go-techblog/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echoインスタンスを作成
	e := echo.New()

	// Middlewareの実行
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRF())

	// src/css ディレクトリ配下のファイルに css のパスでアクセス可能にする
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	// DBの設定
	db := db.ConnectDB(e)
	defer db.Close()
	repository.SetDB(db)

	// ルーティング
	router.CreateMux(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// サーバ起動
	e.Logger.Fatal(e.Start(":" + port))
}
