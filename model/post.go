package model

import (
	"fmt"
	"time"
)

type Post struct {
	ID int `json:"id" gorm:"primaly_key"`
	UserId int `json:"user_id"`
	Title string `json:"title"`
	Detail string `json:"detail"`
	Limit string `json:"limit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User User `json:user gorm:"foreignkey:UserId"`
	Tasks []Task `json:"tasks" gorm:"foreignkey:PostId"`
	Comments []Comment `json:"comments" gorm:"foreignkey:PostId"`
	Favorites []Favorite `json:"favorites" gorm:"foreignkey:PostId"`
}

type Posts []Post

func CreatePost(post *Post) {
	db := Init()
	db.Create(post)
}

func FindAllPosts() Posts {
	var posts Posts
	db := Init()
	db.Preload("Tasks").Preload("Favorites").Preload("Comments").Order("created_at desc").Find(&posts)
	return posts
}

func FindAllPostRanking() Posts {
	var posts Posts
	db := Init()
	db.Preload("Tasks").Preload("Favorites").Preload("Comments").Find(&posts)
	return posts
}

func FindPosts(p *Post) Posts {
	var posts Posts
	db := Init()
	db.Preload("Tasks").Preload("Favorites").Preload("Comments").Where(p).Order("created_at desc").Find(&posts)
	return posts
}

func FindPost(p *Post) Post {
	var post Post
	db := Init()
	db.Preload("Tasks").Preload("Favorites").Preload("Comments").Where(p).Order("created_at desc").Find(&post)
	return post
}

func DeletePost(p *Post) error {
	var posts Posts
	db := Init()
	if rows := db.Where(p).Delete(&posts).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Post (%v) to delete", p)
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

