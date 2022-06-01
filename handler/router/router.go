package router

import (
	"github.com/HarukiIdo/go-techblog/handler"
	"github.com/HarukiIdo/go-techblog/middle"
	"github.com/HarukiIdo/go-techblog/repository"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// NewRouter ...
func NewRouter(e *echo.Echo, articleDB *sqlx.DB) {

	// Basic認証機能を提供するMiddleware
	auth := e.Group("", middle.BasicAuth())

	newRepository := repository.NewArticleRepository(articleDB)
	newHandler := handler.NewArticleHandler(*newRepository)

	// ルーティングの設定
	// TOPページに記事の一覧を表示
	e.GET("/", newHandler.ArticleIndex)

	// Basic認証により認証済みの場合の処理
	auth.GET("/auth", newHandler.ArticleIndex)
	auth.GET("/auth/articles/:articleID", newHandler.ArticleShow)

	// 記事に関連する画面を返す処理
	e.GET("/articles", newHandler.ArticleIndex)                // 一覧画面
	e.GET("/articles/new", newHandler.ArticleNew)              // 新規作成画面
	e.GET("/articles/:articleID", newHandler.ArticleShow)      // 詳細画面
	e.GET("/articles/:articleID/edit", newHandler.ArticleEdit) // 編集画面

	// JSONを返却する処理
	e.GET("/api/articles", newHandler.ArticleList)
	e.POST("/api/articles", newHandler.ArticleCreate)
	e.DELETE("/api/articles/:articleID", newHandler.ArticleDelete)
	e.PATCH("/api/articles/:articleID", newHandler.ArticleUpdate)
}
