package route

import (
	"../api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェアはここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/signup", api.SignupPage())
	e.POST("/signup", api.Signup)
	e.POST("/login", api.Login)
	// TODO ログアウト機能
	//おまじない
	return e
}