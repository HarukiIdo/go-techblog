package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HarukiIdo/go-techblog/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestArticleCreate(t *testing.T) {

}

func TestArticleIndex(t *testing.T) {

	// Echoインスタンスを作成
	e := echo.New()

	// Getリクエストとレスポンスを作成
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	// 新しいコンテキストを作成
	c := e.NewContext(req, rec)

	// FIXME: いらないかも
	//c.SetPath("/articles")

	// DBのモックを作成
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("sqlxmock.Newx() failure %s", err)
	}
	defer db.Close()
	rows := sqlxmock.NewRows([]string{"id"})
	mock.ExpectQuery("SELECT * FROM articles WHERE id < $").WillReturnRows(rows)

	r := repository.NewArticleRepository(db)
	h := NewArticleHandler(r)

	err = h.ArticleIndex(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, mock, rec.Body.String())
	}

}

func TestArticleUpdate(t *testing.T) {

}

func TestArticleDelete(t *testing.T) {

}

func TestArticleEdit(t *testing.T) {

}
