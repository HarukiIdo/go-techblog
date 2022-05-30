package router

import (
	"github.com/HarukiIdo/go-techblog/handler"
	"github.com/HarukiIdo/go-techblog/middle"
	"github.com/labstack/echo/v4"
)

// CreateMux ...
func CreateMux(e *echo.Echo) {

	// Basic認証機能を提供するMiddleware
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
}
