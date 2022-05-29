package _test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/HarukiIdo/go-techblog/model"
)

var (
	var articles model.Article
	mocDB = map[string]*articles{

	}

)

func TestArticleCreate(t *testing.T) {

}

func TestArticleIndex(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/articles")
	c.set
}

func TestArticleUpdate(t *testing.T) {

}

func TestArticleDelete(t *testing.T) {

}

func TestArticleEdit(t *testing.T) {

}
