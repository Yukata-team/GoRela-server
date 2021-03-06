package model

import "time"

type Comment struct {
	ID int `json:"id" gorm:"primaly_key"`
	PostId int `json:"post_id"`
	UserId int `json:"user_id"`
	Content string `json:"content"`
	User User `json:"user" gorm:"foreignkey:UserId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comments []Comment

func CreateComment (comment *Comment) {
	db := Init()
	db.Create(comment)
}