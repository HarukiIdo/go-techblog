package router

import (
	"github.com/HarukiIdo/go-techblog/di"
	"github.com/HarukiIdo/go-techblog/middle"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewRouter ...
func NewRouter(e *echo.Echo, db *sqlx.DB) {

	// リポジトリとハンドラをDIする
	ah := di.InitArticle(db)

	// Basic認証により認証済みの場合の処理
	auth := e.Group("", middle.BasicAuth())
	auth.GET("/auth", ah.ArticleIndex)
	auth.GET("/auth/articles/:articleID", ah.ArticleShow)

	// ルーティングの設定
	// TOPページに記事の一覧を表示
	e.GET("/", ah.ArticleIndex)

	// 記事に関連する画面を返す処理
	e.GET("/articles", ah.ArticleIndex)
	e.GET("/articles/new", ah.ArticleNew)
	e.GET("/articles/:articleID", ah.ArticleShow)
	e.GET("/articles/:articleID/edit", ah.ArticleEdit)

	// JSONを返却する処理
	e.GET("/api/articles", ah.ArticleList)
	e.POST("/api/articles", ah.ArticleCreate)
	e.DELETE("/api/articles/:articleID", ah.ArticleDelete)
	e.PATCH("/api/articles/:articleID", ah.ArticleUpdate)

	// Middlewareの呼び出し
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRF())
}
