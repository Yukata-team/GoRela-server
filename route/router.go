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

	posts := e.Group("/posts")
	//api下はJWTの認証が必要
	posts.Use(middleware.JWTWithConfig(handler.Config))
	posts.GET("", handler.GetPosts)
	posts.POST("", handler.AddPost)
	posts.DELETE("/:id", handler.DeletePost)
	posts.PUT("/:id", handler.UpdatePost)

	return e
}