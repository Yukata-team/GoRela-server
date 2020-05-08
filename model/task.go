package model

import "fmt"

type Task struct {
	ID int `json:"id" gorm:"primaly_key"`
	PostId int `json:"post_id"`
	Content string `json:"content"`
	Isdone bool `json:"isdone"`
}

func FindTask(t *Task) Task {
	var task Task
	db := Init()
	db.Where(t).Find(&task)
	return task
}

func UpdateTask(t *Task) error {
	rows := db.Model(t).Update(map[string]interface{}{
		"isdone": t.Isdone,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Task (%v) to update", t)
	}
	return nil
}
