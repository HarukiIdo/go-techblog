package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/HarukiIdo/go-techblog/model"
	"github.com/HarukiIdo/go-techblog/repository"
	"github.com/labstack/echo/v4"
)

// ArticleCreateOutput ...
type ArticleCreateOutput struct {
	Article          *model.Article
	Message          string
	ValidationErrors []string
}

// ArticleCreate ...
func AriticleCreate(c echo.Context) error {

	// フォームの内容を格納する構造体とレスポンスとして返却する構造体を宣言
	var article model.Article
	var output ArticleCreateOutput

	// フォームの内容を構造体に埋め込む
	if err := c.Bind(&article); err != nil {
		c.Logger().Error(err.Error)
		return c.JSON(http.StatusBadRequest, output)
	}

	// repositoryを呼び出して、フォーム内容の保存処理を実行
	res, err := repository.ArticleCreate(&article)
	if err != nil {
		c.Logger().Error(err.Error)
		return c.JSON(http.StatusInternalServerError, output)
	}

	// SQL実行結果から作成されたレコードのIDを取得
	id, _ := res.LastInsertId()
	article.ID = int(id)

	output.Article = &article

	return c.JSON(http.StatusOK, output)
}

// テンプレートエンジンに埋め込む記事の中身となるデータを渡す
func Articleindex(c echo.Context) error {
	//記事データの一覧を取得する
	ariticles, err := repository.ArticleList()
	if err != nil {
		log.Println(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// dataに取得した記事データを格納する
	data := map[string]interface{}{
		"Message":  " Article Index",
		"Now":      time.Now(),
		"Articles": ariticles,
	}
	return render(c, "article/index.html", data)
}

// ArticleNew ...
func ArticleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}
	return render(c, "article/new.html", data)
}

// ArticleShow ...
func ArticleShow(c echo.Context) error {

	//urlのパスからパラメータを取得
	id, _ := strconv.Atoi(c.Param("id"))
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
		"ID":      id,
	}
	return render(c, "article/show.html", data)
}

// AritcleEdit ...
func ArticleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data := map[string]interface{}{
		"Message": "Article Edit",
		"Now":     time.Now(),
		"ID":      id,
	}
	return render(c, "article/edit.html", data)
}
