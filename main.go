package main

import (
	"log"
	"os"

	"github.com/HarukiIdo/go-techblog/db"
	"github.com/HarukiIdo/go-techblog/handler"
	"github.com/HarukiIdo/go-techblog/handler/router"
	"github.com/HarukiIdo/go-techblog/middle"
	"github.com/HarukiIdo/go-techblog/repository"
	"github.com/joho/godotenv"
)

func main() {
	// 環境変数の読み込み
	loadEnv()

	var e = router.CreateMux()
	db := db.ConnectDB(e)
	defer db.Close()
	repository.SetDB(db)
	auth := e.Group("", middle.BasicAuth())

	// ルーティングの設定
	// TOPページに記事の一覧を表示
	e.GET("/", handler.ArticleIndex)
	auth.GET("/auth", handler.ArticleIndex)
	auth.GET("/auth/articles/:articleID", handler.ArticleShow)

	// 記事に関連する画面を返す処理
	e.GET("/articles", handler.ArticleIndex)                // 一覧画面
	e.GET("/articles/new", handler.ArticleNew)              // 新規作成画面
	e.GET("/articles/:articleID", handler.ArticleShow)      // 詳細画面
	e.GET("/articles/:articleID/edit", handler.ArticleEdit) // 編集画面

	// JSONを返却する処理
	e.GET("/api/articles", handler.ArticleList)
	e.POST("/api/articles", handler.ArticleCreate)
	e.DELETE("/api/articles/:articleID", handler.ArticleDelete)
	e.PATCH("/api/articles/:articleID", handler.ArticleUpdate)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".envファイルの読み込みに失敗しました")
	}
}
