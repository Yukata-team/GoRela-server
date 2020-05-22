package model

import "fmt"

type Relation struct {
	ID int `json:"id" gorm:"primaly_key"`
	FollowUserId int `json:"follow_user_id"`
	FollowedUserId int `json:"followed_user_id"`
}

func CreateRelation (r *Relation) {
	db := Init()
	db.Create(r)
}

func DeleteRelation (r *Relation) error {
	var relation Relation
	db := Init()
	if rows := db.Where(r).Delete(&relation).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Favo (%v) to delete", r)
	}
	return nil
}