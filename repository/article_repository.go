package repository

import "github.com/HarukiIdo/go-techblog/model"

// ArticleList ...
func ArticleList() ([]*model.Article, error) {
	query := `SELECT * FROM articles;`

	var articles []*model.Article

	if err := db.Select(&articles, query); err != nil {
		return nil, err
	}
	return articles, nil
}
