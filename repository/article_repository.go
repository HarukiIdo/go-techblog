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

	// 引数で渡されたカーソルの値が0以下の場合は、int型の最大値で置き換える
	if cursor <= 0 {
		cursor = math.MaxInt32
	}

	query := `SELECT * FROM articles WHERE id < ? ORDER BY id desc LIMIT 10`

	// クエリ結果を格納するためのスライスを初期化
	articles := make([]*model.Article, 0, 10)

	// クエリを実行
	//  Selectは複数レコードを取得することが可能
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

	query := "INSERT INTO articles(title, body, createdat, updatedat) VALUES (:title, :body, :createdat, :updatedat);"

	// トランザクションを開始
	tx := db.MustBegin()

	// SQLを実行
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

// ArticleDelete ...
func ArticleDelete(id int) error {

	query := "DELETE FROM articles WHERE id = ?;"

	// トランザクションを開始
	tx := db.MustBegin()

	// SQLを実行
	// エラーが発生した場合は、ロールバックしエラー内容を返す
	if _, err := tx.Exec(query, id); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// ArticleGetByID ...
func ArticleGetByID(id int) (*model.Article, error) {

	query := "SELECT * FROM articles WHERE id = ?;"

	// クエリ結果を格納する変数
	var article model.Article

	// SQLを実行
	// エラーが発生した場合はエラーを返却
	if err := db.Get(&article, query, id); err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("ID:%d", article.ID)

	return &article, nil
}
