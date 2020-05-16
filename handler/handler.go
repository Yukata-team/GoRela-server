package handler

import (
	"fmt"
	"github.com/Yukata-team/GoRela-server/model"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// Post
func AddPost(c echo.Context) error {
	post := new(model.Post)
	if err := c.Bind(post); err != nil {
		return err
	}

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
	posts := model.FindAllPosts()

	return c.JSON(http.StatusOK, posts)
}

//func GetUserPosts(c echo.Context) error {
//	userId := userIDFromToken(c)
//	if user := model.FindUser(&model.User{ID: userId}); user.ID == 0 {
//		return echo.ErrNotFound
//	}
//
//	posts := model.FindPosts(&model.Post{UserId: userId})
//
//	return c.JSON(http.StatusOK, posts)
//}

func ShowPost(c echo.Context) error {
	userId := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	//指定されたURL上のIDが数字か
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	posts := model.FindPost(&model.Post{ID: postID})

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

func UpdatePost(c echo.Context) error  {
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

	fmt.Println(1)
	for i, task := range post.Tasks {
		npost.Tasks[i].ID = task.ID
		pp.Println(npost.Tasks[i])
	}
	post.Title = npost.Title
	post.Detail = npost.Detail
	post.Tasks = npost.Tasks

	fmt.Println(2)
	pp.Println(post)

	if err := model.UpdatePost(&post); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func AddComment(c echo.Context) error {

	comment := new(model.Comment)
	if err := c.Bind(comment); err != nil {
		return err
	}

	if comment.Content == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	userId := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	//指定されたURL上のIDが数字か
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	//ユーザーが作成した該当IDのPostがデータベース上に存在するか
	posts := model.FindPosts(&model.Post{ID: postId})
	if len(posts) == 0 {
		return echo.ErrNotFound
	}
	post := posts[0]

	comment.UserId = userId
	comment.PostId = post.ID
	model.CreateComment(comment)

	return c.JSON(http.StatusCreated, comment)
}

func AddFavo(c echo.Context) error {
	favo := new(model.Favorite)
	if err := c.Bind(favo); err != nil {
		return err
	}

	userId := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	favo.UserId = userId
	favo.PostId = postId
	model.CreateFavo(favo)

	return c.JSON(http.StatusCreated, favo)
}

func DeleteFavo(c echo.Context) error {
	userId := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.DeleteFavo(&model.Favorite{PostId: postId, UserId: userId}); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func GetUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	user := model.FindUser(&model.User{ID: userID})
	if user.ID == 0 {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, user)
}