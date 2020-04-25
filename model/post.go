package model

type Post struct {
	ID int `json:"id" gorm:"primaly_key"`
	UID int `json:"user_id"`
	Title string `json:"title"`
	Detail string `json:"detail"`
}

func CreatePost (post *Post) {
	db := Init()
	db.Debug().Create(post)
}