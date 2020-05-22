package handler

import (
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
	if user := model.FindUserOnly(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	post.UserId = userId
	pp.Println(post)
	model.CreatePost(post)

	return c.JSON(http.StatusCreated, post)
}

func GetPosts(c echo.Context) error {
	posts := model.FindAllPosts()

	for i:=0; i<len(posts); i++ {
		posts[i].User = model.FindUserOnly(&model.User{ID: posts[i].UserId})
	}

	return c.JSON(http.StatusOK, posts)
}

func ShowPost(c echo.Context) error {
	userId := userIDFromToken(c)
	if user := model.FindUserOnly(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	//指定されたURL上のIDが数字か
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	post := model.FindPost(&model.Post{ID: postID})
	post.User = model.FindUserOnly(&model.User{ID: post.UserId})

	for i:=0; i<len(post.Comments); i++ {
		post.Comments[i].User = model.FindUserOnly(&model.User{ID: post.Comments[i].UserId})
	}

	return c.JSON(http.StatusOK, post)
}


func DeletePost(c echo.Context) error {
	userId := userIDFromToken(c)

	//受け取ったJWT内のユーザーIDがデータベースに存在するか
	if user := model.FindUserOnly(&model.User{ID: userId}); user.ID == 0 {
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
	if user := model.FindUserOnly(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	//指定されたURL上のIDが数字か
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	//ユーザーが作成した該当IDのPostがデータベース上に存在するか
	post := model.FindPost(&model.Post{ID: postID, UserId: userId})

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

	for i, task := range post.Tasks {
		npost.Tasks[i].ID = task.ID
		pp.Println(npost.Tasks[i])
	}
	post.Title = npost.Title
	post.Detail = npost.Detail
	post.Tasks = npost.Tasks

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
	if user := model.FindUserOnly(&model.User{ID: userId}); user.ID == 0 {
		return echo.ErrNotFound
	}

	//指定されたURL上のIDが数字か
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	//ユーザーが作成した該当IDのPostがデータベース上に存在するか
	post := model.FindPost(&model.Post{ID: postId})

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
	if user := model.FindUserOnly(&model.User{ID: userId}); user.ID == 0 {
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
	if user := model.FindUserOnly(&model.User{ID: userId}); user.ID == 0 {
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

//func GetPostRanking(c echo.Context) error {
//	posts := model.FindAllPosts()
//
//	for i:=0; i<len(posts); i++ {
//		posts[i].User = model.FindUserOnly(&model.User{ID: posts[i].UserId})
//		posts[i].FavoCounts = model.CountFavo(&model.Favorite{PostId: posts[i].ID})
//	}
//
//	posts = model.Preload().
//
//	return c.JSON(http.StatusOK, posts)
//}

// User

func GetMyPage(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	user := model.FindUser(&model.User{ID: userID})
	if user.ID == 0 {
		return echo.ErrNotFound
	}
	user.Posts = model.FindPosts(&model.Post{UserId: userID})

	return c.JSON(http.StatusOK, user)
}

func GetUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	user := model.FindUserOnly(&model.User{ID: userID})
	if user.ID == 0 {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	userId := userIDFromToken(c)

	//受け取ったJWT内のユーザーIDがデータベースに存在するか
	user := model.FindUserOnly(&model.User{ID: userId})
	if user.ID == 0 {
		return echo.ErrNotFound
	}

	//ユーザーがリクエストしたpost
	nuser := new(model.User)
	if err := c.Bind(nuser); err != nil {
		return err
	}

	if nuser.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "NoName",
		}
	}

	user.Name = nuser.Name
	user.Introduction = nuser.Introduction

	if err := model.UpdateUser(&user); err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateTask(c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	task := model.FindTask(&model.Task{ID: taskID})
	if task.ID == 0 {
		return echo.ErrNotFound
	}

	task.IsDone = !task.IsDone

	pp.Println(task)

	if err := model.UpdateTask(&task); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func AddRelation(c echo.Context) error {
	relation := new(model.Relation)
	if err := c.Bind(relation); err != nil {
		return err
	}

	followID := userIDFromToken(c)
	if user := model.FindUserOnly(&model.User{ID: followID}); user.ID == 0 {
		return echo.ErrNotFound
	}

	followedID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if followID == followedID {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "type:FollowingMyself",
		}
	}

	relation.FollowUserId = followID
	relation.FollowedUserId = followedID
	model.CreateRelation(relation)

	return c.JSON(http.StatusCreated, relation)
}

func DeleteRelation(c echo.Context) error {
	followID := userIDFromToken(c)
	if user := model.FindUserOnly(&model.User{ID: followID}); user.ID == 0 {
		return echo.ErrNotFound
	}

	followedID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.DeleteRelation(&model.Relation{FollowUserId: followID, FollowedUserId: followedID}); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}