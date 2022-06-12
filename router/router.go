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

	// ルーティングの設定
	// TOPページに記事の一覧を表示
	e.GET("/", ah.ArticleIndex)

	// Basic認証により認証済みの場合の処理
	auth := e.Group("", middle.BasicAuth())
	auth.GET("/auth", ah.ArticleIndex)
	auth.GET("/auth/articles/:articleID", ah.ArticleShow)

	// 記事に関連する画面を返す処理
	e.GET("/articles", ah.ArticleIndex)                // 一覧画面
	e.GET("/articles/new", ah.ArticleNew)              // 新規作成画面
	e.GET("/articles/:articleID", ah.ArticleShow)      // 詳細画面
	e.GET("/articles/:articleID/edit", ah.ArticleEdit) // 編集画面

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
