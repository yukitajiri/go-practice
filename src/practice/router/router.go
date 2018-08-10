package router

import (
	"github.com/labstack/echo"
	"net/http"
)

func Build () *echo.Echo {
	e := echo.New()

	e.GET("/", hello)
	return e
}


//　レスポンスを書き込む先とクライアントからのリクエストを受け取る
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

