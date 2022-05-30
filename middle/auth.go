package middle

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// BasicAuth ...
func BasicAuth() echo.MiddlewareFunc {

	authUser := os.Getenv("AUTH_USER")
	authPassword := os.Getenv("AUTH_PASSWORD")

	// 認証成功か失敗を判定する関数
	basicAuthValidator := func(username, password string, c echo.Context) (bool, error) {
		if username == authUser && password == authPassword {
			return true, nil
		}
		return false, nil
	}

	return middleware.BasicAuth(basicAuthValidator)
}
