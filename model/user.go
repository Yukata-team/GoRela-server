package model

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Posts []Post `json:"posts" gorm:"foreignkey:UserId"`
}

func CreateUser(user *User) {
	db := Init()
	db.Debug().Create(user)
}

func FindUser(u *User) User {
	var user User
	db := Init()
	db.Preload("Posts").Where(u).First(&user)
	return user
}
