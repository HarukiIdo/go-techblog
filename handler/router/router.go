package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CreateMux ... 
func CreateMux() *echo.Echo {

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

	return e
}
