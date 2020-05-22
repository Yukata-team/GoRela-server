package model

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name 	 string `json:"name"`
	Introduction string `json:"introduction"`
	Posts 	[]Post 	`json:"posts" gorm:"foreignkey:UserId"`
	Follows []Relation `json:"follows" gorm:"foreignkey:FollowUserId"`
	Followers []Relation `json:"followers" gorm:"foreignkey:FollowedUserId"`
}

func CreateUser(user *User) {
	db := Init()
	db.Debug().Create(user)
}

func FindUser(u *User) User {
	var user User
	db := Init()
	db.Preload("Follows").Preload("Followers").Preload("Posts").Preload("Posts.Tasks").Preload("Posts.Favorites").Preload("Posts.Comments").Where(u).First(&user)
	return user
}

func FindUserOnly(u *User) User {
	var user User
	db := Init()
	db.Where(u).First(&user)
	return user
}
