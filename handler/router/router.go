package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ルーティング
func CreateMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	// src/css ディレクトリ配下のファイルに css のパスでアクセス可能にする
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	return e
}
