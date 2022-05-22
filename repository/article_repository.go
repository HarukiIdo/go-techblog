package repository

import (
	"database/sql"
	"log"
	"math"
	"time"

	"github.com/HarukiIdo/go-techblog/model"
)

// ArticleListByCursor ...
func ArticleListByCursor(cursor int) ([]*model.Article, error) {

	if cursor <= 0 {
		cursor = math.MaxInt32
	}

	query := `SELECT * FROM articles WHERE id < ? ORDER BY id desc LIMIT 10`

	articles := make([]*model.Article, 0, 10)

	if err := db.Select(&articles, query, cursor); err != nil {
		return nil, err
	}
	return articles, nil
}

// ArticleCreate ...
func ArticleCreate(article *model.Article) (sql.Result, error) {

	//現在時刻を取得
	now := time.Now()
	article.CreatedAt = now
	article.UpdatedAt = now

	query := `INSERT INTO articles(title, body, createdat, updatedat) VALUES (:title, :body, :createdat, :updatedat);`

	// トランザクションを開始
	tx := db.MustBegin()

	// SQLを実行する
	// エラーが発生した場合はロールバック
	res, err := tx.NamedExec(query, article)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, err
	}
	// SQLの実行に成功したらコミット
	tx.Commit()

	// SQLの実行結果を返す
	return res, nil
}
