package repository

import (
	"database/sql"
	"log"
	"math"
	"time"

	"github.com/HarukiIdo/go-techblog/model"
	"github.com/jmoiron/sqlx"
)

type ArticleRepository interface {
	ArticleListByCursor(cusor int) ([]*model.Article, error)
	ArticleCreate(article *model.Article) (sql.Result, error)
	ArticleUpdate(article *model.Article) (sql.Result, error)
	ArticleGetByID(id int) (*model.Article, error)
	ArticleDelete(id int) error
}

type articleRepository struct {
	db *sqlx.DB
}

// NewArticleRepository returns new articleRepository
func NewArticleRepository(db *sqlx.DB) ArticleRepository {
	return &articleRepository{db}
}

// ArticleListByCursor ...
func (r *articleRepository) ArticleListByCursor(cursor int) ([]*model.Article, error) {

	// 引数で渡されたカーソルの値が0以下の場合は、int型の最大値で置き換える
	if cursor <= 0 {
		cursor = math.MaxInt32
	}

	query := `SELECT * FROM articles WHERE id < ? ORDER BY id desc LIMIT 10`

	// クエリ結果を格納するためのスライスを初期化
	articles := make([]*model.Article, 0, 10)

	// クエリを実行
	//  Selectは複数レコードを取得することが可能
	if err := r.db.Select(&articles, query, cursor); err != nil {
		log.Println(err)
		return nil, err
	}
	return articles, nil
}

// ArticleCreate ...
func (r *articleRepository) ArticleCreate(article *model.Article) (sql.Result, error) {

	//現在時刻を取得
	now := time.Now()
	article.CreatedAt = now
	article.UpdatedAt = now

	query := "INSERT INTO articles(title, body, createdat, updatedat) VALUES (:title, :body, :createdat, :updatedat);"

	// トランザクションを開始
	tx := r.db.MustBegin()

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

func (r *articleRepository) ArticleUpdate(article *model.Article) (sql.Result, error) {

	// 現在日時を記事構造体に設定
	article.UpdatedAt = time.Now()

	query := "UPDATE articles SET title = :title, body = :body, updatedat := updatedat where id = :id;"

	// トランザクションを開始
	tx := r.db.MustBegin()

	// SQLを実行
	// クエリ文字列内の:title, :body, :updatedatには、
	// 第２引数のArticle構造体のTitle, Body, IDがbindされる
	res, err := tx.NamedExec(query, article)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// エラーがない場合はコミットし、
	// SQLの実行結果を返す
	tx.Commit()
	return res, nil
}

// ArticleGetByID ...
func (r *articleRepository) ArticleGetByID(id int) (*model.Article, error) {

	query := "SELECT * FROM articles WHERE id = ?;"

	// クエリ結果を格納する変数
	var article model.Article

	// SQLを実行
	// エラーが発生した場合はエラーを返却
	if err := r.db.Get(&article, query, id); err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("ID:%d", article.ID)

	return &article, nil
}

// ArticleDelete ...
func (r *articleRepository) ArticleDelete(id int) error {

	query := "DELETE FROM articles WHERE id = ?;"

	// トランザクションを開始
	tx := r.db.MustBegin()

	// SQLを実行
	// エラーが発生した場合は、ロールバックしエラー内容を返す
	if _, err := tx.Exec(query, id); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
