package model

type Relation struct {
	ID int `json:"id" gorm:"primaly_key"`
	FollowUserId int `json:"follow_user_id"`
	FollowedUserId int `json:"followed_user_id"`
}

func CreateRelation (r *Relation) {
	db := Init()
	db.Create(r)
}

//func DeleteRelation (r *Relation) error {
//	var follow Follow
//	db := Init()
//	if rows := db.Delete(&follow).RowsAffected; rows == 0 {
//		return fmt.Errorf("Could not find Favo (%v) to delete", f)
//	}
//	return nil
//}