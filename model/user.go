package model

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) {
	db := Init()
	db.Debug().Create(user)
}

func FindUser(u *User) User {
	var user User
	db := Init()
	db.Where(u).First(&user)
	return user
}
