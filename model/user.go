package model

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) {
	dbinstance := Init()
	dbinstance.Debug().Create(user)
	//db.Create(user)
}

// func FindUser(u *User) User {
// 	var user User
// 	db.Where(u).First(&user)
// 	return user
// }
