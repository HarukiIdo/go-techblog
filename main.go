package main

import (
	"log"
	"os"

	"github.com/HarukiIdo/go-techblog/handler"
	"github.com/HarukiIdo/go-techblog/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sqlx.DB
var e = createMux()

func main() {
	db = connectDB()
	repository.SetDB(db)

	//ルーティングの設定
	e.GET("/", handler.Articleindex)
	e.GET("/new", handler.ArticleNew)
	e.GET(":id", handler.ArticleShow)
	e.GET(":id/edit", handler.ArticleEdit)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

// ルーティング
func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	// src/css ディレクトリ配下のファイルに css のパスでアクセス可能にする
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	return e
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	driver := os.Getenv("DRIVER")
	log.Println(dsn)
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}
