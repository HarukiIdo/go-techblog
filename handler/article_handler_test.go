package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestArticleCreate(t *testing.T) {

}

func TestArticleIndex(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/articles")
	//h := handler.ArticleIndex(e)
}

func TestArticleUpdate(t *testing.T) {

}

func TestArticleDelete(t *testing.T) {

}

func TestArticleEdit(t *testing.T) {

}
