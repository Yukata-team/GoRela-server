package handler

import (
	"github.com/Yukata-team/GoRela-server/model"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"net/http"
)

func AddPost(c echo.Context) error {
	post := new(model.Post)
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

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	post.UID = uid
	model.CreatePost(post)

	return c.JSON(http.StatusCreated, post)
}