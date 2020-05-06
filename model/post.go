package model

import (
	"fmt"
	"github.com/k0kubun/pp"
)

type Task struct {
	ID int `json:"id" gorm:"primaly_key"`
	PostId int `json:"post_id"`
	Content string `json:"content"`
	Isdone bool `json:"isdone"`
}

type Post struct {
	ID int `json:"id" gorm:"primaly_key"`
	UserId int `json:"user_id"`
	Title string `json:"title"`
	Detail string `json:"detail"`
	Tasks []Task `json:"tasks" gorm:"foreignkey:PostId"`
}

type Posts []Post

func CreatePost (post *Post) {
	db := Init()
	db.Create(post)
}

func FindPosts(p *Post) Posts {
	var posts Posts
	db := Init()
	db.Preload("Tasks").Where(p).Find(&posts)
	return posts
}

func DeletePost(p *Post) error {
	var posts Posts
	db := Init()
	db.Preload("Tasks").Where(p).Find(&posts)
	//if rows := db.Delete(&posts).RowsAffected; rows == 0 {
	if rows := db.Table("posts").Select("posts.*").Joins("left join tasks on tasks.post_id = posts.id").Delete(&posts).RowsAffected; rows == 0 {
		return fmt.Errorf("Coule not find Post (%v) to delete", p)
	}
	return nil
}

func UpdatePost(p *Post) error {
	db := Init()
	rows := db.Model(p).Update(map[string]interface{}{
		"title": p.Title,
		"detail": p.Detail,
		//"tasks": p.Tasks,
	}).RowsAffected
	pp.Println(p)
	if rows == 0 {
		return fmt.Errorf("Could not find Post (%v) to update", p)
	}
	return nil
}

