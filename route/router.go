package route

import (
	"github.com/Yukata-team/GoRela-server/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	// Echoのインスタンス作る
	e := echo.New()

	// CORSの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// 全てのリクエストで差し込みたいミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/signup", handler.SignupPage())
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)

	posts := e.Group("/posts")
	// posts下はJWTの認証が必要
	posts.Use(middleware.JWTWithConfig(handler.Config))
	posts.GET("", handler.GetPosts)
	posts.POST("", handler.AddPost)
	posts.GET("/:id", handler.ShowPost)
	posts.DELETE("/:id", handler.DeletePost)
	posts.PUT("/:id", handler.UpdatePost)
	posts.POST("/:id/comment", handler.AddComment)
	posts.POST("/:id/favorite", handler.AddFavo)
	posts.DELETE("/:id/favorite", handler.DeleteFavo)

	tasks := e.Group("/tasks")
	// tasks下はJWTの認証が必要
	tasks.PUT("/:id", handler.UpdateTask)

	users := e.Group("/users")
	// users下はJWTの認証が必要
	users.Use(middleware.JWTWithConfig(handler.Config))
	users.GET("/:id/edit", handler.GetUser)
	users.GET("/:id", handler.GetMyPage)
	users.PUT("/:id", handler.UpdateUser)
	users.POST("/:id/follow", handler.AddRelation)
	users.DELETE("/:id/follow", handler.DeleteRelation)

	return e
}