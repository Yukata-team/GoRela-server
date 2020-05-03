package handler

import (
	"github.com/Yukata-team/GoRela-server/model"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func AddPost(c echo.Context) error {
	post := new(model.Post)
	pp.Println(post)
	if err := c.Bind(post); err != nil {
		return err
	}
	pp.Println(post)

	if post.Title == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	userId := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	post.UserId = userId
	model.CreatePost(post)

	return c.JSON(http.StatusCreated, post)
}

func GetPosts(c echo.Context) error {
	userId := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	posts := model.FindPosts(&model.Post{UserId: userId})
	return c.JSON(http.StatusOK, posts)
}

func DeletePost(c echo.Context) error {
	userId := userIDFromToken(c)

	//受け取ったJWT内のユーザーIDがデータベースに存在するか
	if user := model.FindUser(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	//指定されたURL上のIDが数字か
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	//ユーザーが作成したPostがデータベース上に存在するか
	//削除処理
	if err := model.DeletePost(&model.Post{ID: postID, UserId: userId}); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdatePost(c echo.Context) error {
	userId := userIDFromToken(c)

	//受け取ったJWT内のユーザーIDがデータベースに存在するか
	if user := model.FindUser(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	//指定されたURL上のIDが数字か
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	//ユーザーが作成した該当IDのPostがデータベース上に存在するか
	posts := model.FindPosts(&model.Post{ID: postID, UserId: userId})
	if len(posts) == 0 {
		return echo.ErrNotFound
	}
	post := posts[0]

	//ユーザーがリクエストしたpost
	npost := new(model.Post)
	if err := c.Bind(npost); err != nil {
		return err
	}

	if npost.Title == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	post.Title = npost.Title
	post.Detail = npost.Detail

	if err := model.UpdatePost(&post); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}