package handler

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

// テンプレートファイルを配置するディレクトリの相対パス
const tmplPath = "src/templates/"

// pongo2テンプレートエンジンでHTMLを解析、新しいテンプレートを生成する
func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	baseTemplete := pongo2.Must(pongo2.FromCache(tmplPath + file))
	by, err := baseTemplete.ExecuteBytes(data)
	return by, err
}

// テンプレートエンジンを呼び出し、レスポンスと記事の中身を返す
func render(c echo.Context, file string, data map[string]interface{}) error {

	// 定義したhtmlBlob()関数を呼び出し、生成されたHTMLをバイトデータとして受け取る
	b, err := htmlBlob(file, data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	//ステータスコード200でHTMLデータをレスポンス
	return c.HTMLBlob(http.StatusOK, b)
}
