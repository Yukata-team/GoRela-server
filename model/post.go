package model

import "fmt"

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
	Tasks []Task
}

type Posts []Post

func CreatePost (post *Post) {
	db := Init()
	db.Create(post)
}

func FindPosts(p *Post) Posts {
	var posts Posts
	db := Init()
	db.Where(p).Find(&posts)
	return posts
}

func DeletePost(p *Post) error {
	db := Init()
	if rows := db.Where(p).Delete(&Post{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Coule not find Post (%v) to delete", p)
	}
	return nil
}

func UpdatePost(p *Post) error {
	db := Init()
	rows := db.Model(p).Update(map[string]interface{}{
		"title": p.Title,
		"detail": p.Detail,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Post (%v) to update", p)
	}
	return nil
}

