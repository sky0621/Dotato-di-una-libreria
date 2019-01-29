package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// Routing ...
func Routing(e *echo.Echo) {
	InitValidator()

	http.Handle("/", e)

	// 個人の認証を要するWebAPI用のルート
	authGroup := e.Group("/api/v1")

	// 「お知らせ」機能のルーティング("/notices")
	HandleNotice(authGroup)

	// 「ユーザ」機能のルーティング("/users")
	HandleUser(authGroup)
}
