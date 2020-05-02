package route

import (
	"github.com/Yukata-team/GoRela-server/handler"
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
	e.GET("/signup", handler.SignupPage())
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)

	api := e.Group("/api")
	//api下はJWTの認証が必要
	api.Use(middleware.JWTWithConfig(handler.Config))
	api.GET("/posts", handler.GetPosts)
	api.POST("/posts", handler.AddPost)
	api.DELETE("/posts/:id", handler.DeletePost)
	api.PUT("/posts/:id", handler.UpdatePost)

	return e
}