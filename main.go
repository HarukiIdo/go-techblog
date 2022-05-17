package main

import (
	"os"

	"github.com/HarukiIdo/go-techblog/db"
	"github.com/HarukiIdo/go-techblog/handler"
	"github.com/HarukiIdo/go-techblog/handler/router"
	"github.com/HarukiIdo/go-techblog/repository"
)

func main() {
	var e = router.CreateMux()
	db := db.ConnectDB(e)
	repository.SetDB(db)

	//ルーティングの設定
	e.GET("/", handler.Articleindex)
	e.GET("/new", handler.ArticleNew)
	e.GET(":id", handler.ArticleShow)
	e.GET(":id/edit", handler.ArticleEdit)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
