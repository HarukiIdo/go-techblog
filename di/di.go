package di

import (
	"github.com/HarukiIdo/go-techblog/handler"
	"github.com/HarukiIdo/go-techblog/repository"
	"github.com/jmoiron/sqlx"
)

func InitArticle(db *sqlx.DB) handler.ArticleHandler {

	// リポジトリとハンドラをDIする
	ar := repository.NewArticleRepository(db)
	ah := handler.NewArticleHandler(ar)
	return ah
}
