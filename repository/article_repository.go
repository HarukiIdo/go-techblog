package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/HarukiIdo/go-techblog/model"
)

// ArticleList ...
func ArticleList() ([]*model.Article, error) {
	query := `SELECT * FROM articles;`

	var articles []*model.Article

	if err := db.Select(&articles, query); err != nil {
		return nil, err
	}

	//fmt.Println(articles[len(articles)-1].Title)

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
