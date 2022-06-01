package _test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/HarukiIdo/go-techblog/handler"
	"github.com/HarukiIdo/go-techblog/model"
	"github.com/labstack/echo/v4"
)

var (
	articles *model.Article
	mocDB    = map[string]articles{
		"id":        1,
		"title":     "",
		"Body":      "",
		"CreatedAt": time.Time(),
		"UpdatedAt": time.Time(),
	}
)

func TestArticleCreate(t *testing.T) {

}

func initServer() http.Handler {

}

func TestArticleIndex(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/articles")
	h :=handler.ArticleIndex(e)
}

func TestArticleUpdate(t *testing.T) {

}

func TestArticleDelete(t *testing.T) {

}

func TestArticleEdit(t *testing.T) {

}
