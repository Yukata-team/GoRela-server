package model

type Comment struct {
	ID int `json:"id" gorm:"primaly_key"`
	PostId int `json:"post_id"`
	UserId int `json:"user_id"`
	Content string `json:"content"`
}

type Comments []Comment

func CreateComment (comment *Comment) {
	db := Init()
	db.Create(comment)
}