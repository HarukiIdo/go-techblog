package repository

import (
	"testing"

	"github.com/HarukiIdo/go-techblog/custom_error"
	"github.com/HarukiIdo/go-techblog/model"
	"github.com/jmoiron/sqlx"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestArticle_Create(t *testing.T) {
	t.Run(
		"正常系：エラーなし",
		func(t *testing.T) {
			article := &model.Article{
				Title: "article title",
				Body:  "article body",
			}

			// テストケースを用意
			tests := []struct {
				name    string
				db      *sqlx.DB
				article model.Article
				want    error
				wantErr bool
			}{
				{
					name: "OK",
					db: func() *sqlx.DB {
						// モック用のコネクションを作成
						db, mock, err := sqlxmock.Newx()
						if err != nil {
							t.Fatal("sqlxmock.Newx() failure", err)
						}
						rows := sqlxmock.NewRows([]string{"id"}).AddRow()
						mock.ExpectQuery("INSERT INTO articles").WithArgs(article.Title, article.Body, article.CreatedAt, article.UpdatedAt).WillReturnRows(rows)
						return db
					}(),
					article: model.Article{
						Title: "first_title",
						Body:  "first_body",
					},
					want: &custom_error.CreateError{Msg: "create error"},
				},
				{},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					// テスト対象のリポジトリを作成

					r := NewArticleRepository(tt.db)
					t.Cleanup(func() { tt.db.Close() })
					got, err := r.ArticleCreate(&tt.article)
					if (err != nil) != tt.wantErr {
						t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
						return
					}
					if err == nil && got != &tt.article {
						t.Errorf("get = %v, but want = %v", got, tt.want)
					}
				})
			}
		},
	)
}
