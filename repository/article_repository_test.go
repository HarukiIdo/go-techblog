package repository

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/HarukiIdo/go-techblog/model"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestArticle_Create(t *testing.T) {

	t.Run(
		"正常系：エラーなし",
		func(t *testing.T) {
			article := &model.Article{
				Title: "article title",
				Body:  "article body",
			}

			// モック用のコネクションを作成
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()
			dbx := sqlx.NewDb(db, "mysql")

			// SQLのクエリの引数と戻り値が期待値と一致するか
			mock.ExpectPrepare(`INSERT INTO articles`).ExpectExec().WithArgs(article.Title, article.Body, article.CreatedAt, article.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

			// テスト対象のリポジトリを作成
			r := NewArticleRepository(dbx)

			res, err := r.ArticleCreate(article)
			fmt.Println(res)

			// エラーが発生しないことを期待
			assert.NoError(t, err)

			// 上記で指定した通りにモックが呼ばれることを期待
			assert.NoError(t, mock.ExpectationsWereMet())

		},
	)
}
