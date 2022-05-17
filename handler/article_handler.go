package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/HarukiIdo/go-techblog/repository"
	"github.com/labstack/echo/v4"
)

// テンプレートエンジンに埋め込む記事の中身となるデータを渡す
func Articleindex(c echo.Context) error {
	//記事データの一覧を取得する
	ariticles, err := repository.ArticleList()
	if err != nil {
		log.Println(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

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
