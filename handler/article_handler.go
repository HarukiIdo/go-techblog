package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/HarukiIdo/go-techblog/model"
	"github.com/HarukiIdo/go-techblog/repository"
	"github.com/labstack/echo/v4"
)


// ArticleHandler defines repository interface
type ArticleHandler struct {
	ar repository.ArticleRepository
}

// NewArticleHandler returns ArticleHandler based echo.Handler
func NewArticleHandler(ar repository.ArticleRepository) *ArticleHandler {
	return &ArticleHandler{
		ar: ar,
	}
}

// ArticleCreateOutput ...
type ArticleCreateOutput struct {
	Article          *model.Article
	Message          string
	ValidationErrors []string
}

// ArticleCreate ...
func (h *ArticleHandler) ArticleCreate(c echo.Context) error {

	// フォームの内容を格納する構造体とレスポンスとして返却する構造体を宣言
	var article model.Article
	var output ArticleCreateOutput

	// フォームの内容を構造体に埋め込む
	if err := c.Bind(&article); err != nil {
		c.Logger().Error(err.Error)
		return c.JSON(http.StatusBadRequest, output)
	}

	// repositoryを呼び出して、フォーム内容の保存処理を実行
	res, err := h.ar.ArticleCreate(&article)
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

// ArticleList ...
func (h *ArticleHandler)ArticleList(c echo.Context) error {

	data := map[string]interface{}{}
	return render(c, "article/index.html", data)
}

// ArticleIndex ...
func (h *ArticleHandler) ArticleIndex(c echo.Context) error {
	//記事データの一覧を取得する
	ariticles, err := h.ar.ArticleListByCursor(0)
	if err != nil {
		log.Println(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// dataに取得した記事データを格納する
	data := map[string]interface{}{
		"Articles": ariticles,
	}
	return render(c, "article/index.html", data)
}

// ArticleNew ...
func (h *ArticleHandler) ArticleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}
	return render(c, "article/new.html", data)
}

// ArticleDelete ...
func (h *ArticleHandler) ArticleDelete(c echo.Context) error {

	// パスパラメータから記事IDを取得
	id, _ := strconv.Atoi(c.Param("articleID"))

	// 記事削除処理を呼び出す
	if err := h.ar.ArticleDelete(id); err != nil {
		c.Logger().Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Article %d is deleted", id))
}

// ArticleShow ...
func (h *ArticleHandler) ArticleShow(c echo.Context) error {

	// パスパラメータから記事のIDを取得
	// 文字列を数値型にキャスト
	id, _ := strconv.Atoi(c.Param("articleID"))

	// 記事データを取得
	article, err := h.ar.ArticleGetByID(id)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// テンプレートに渡すデータをmapに格納
	data := map[string]interface{}{
		"Article": article,
	}

	// テンプレートファイルとデータを指定してHTMLを生成し、クライアントに返却
	return render(c, "article/show.html", data)
}

// ArticleEdit ...
func (h *ArticleHandler) ArticleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("articleID"))

	article, err := h.ar.ArticleGetByID(id)

	if err != nil {
		c.Logger().Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// テンプレートに渡すデータをmapに格納
	data := map[string]interface{}{
		"Article": article,
	}
	return render(c, "article/edit.html", data)
}

func (h *ArticleHandler) ArticleUpdate(c echo.Context) error {

	// リクエスト送信元のパスを取得
	// パスから記事IDを抽出
	ref := c.Request().Referer()
	refID := strings.Split(ref, "/")[4]

	// リクエストURLのパスパラメータから記事IDを取得
	reqID := c.Param("articleID")

	// 編集画面で表示している記事と更新する記事が異なる場合
	// 更新処理をせずにエラーを返却
	if refID != reqID {
		log.Println("ID not match")
		return c.JSON(http.StatusBadRequest, "")
	}

	var article model.Article
	var output ArticleCreateOutput

	// フォームで送信されたデータを変数に格納
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, output)
	}

	// IDをint型にキャスト
	articleID, _ := strconv.Atoi(reqID)
	article.ID = articleID

	_, err := h.ar.ArticleUpdate(&article)
	if err != nil {
		c.Logger().Error(err.Error())
		output.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, output)
	}
	output.Article = &article

	return c.JSON(http.StatusOK, output)
}
