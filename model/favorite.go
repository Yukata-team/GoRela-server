package model

import "fmt"

type Favorite struct {
	ID int `json:"id" gorm:"primaly_key"`
	PostId int `json:"post_id"`
	UserId int `json:"user_id"`
}

func CreateFavo (favo *Favorite) {
	db := Init()
	db.Create(favo)
}

func DeleteFavo (f *Favorite) error {
	var favo Favorite
	db := Init()
	if rows := db.Delete(&favo).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Favo (%v) to delete", f)
	}
	return nil
}

//func CountFavo (f *Favorite) {
//	db := Init()
//	db.Where(f).Find(&favo)
//}
