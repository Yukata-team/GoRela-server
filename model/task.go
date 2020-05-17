package model

import "fmt"

type Task struct {
	ID int `json:"id" gorm:"primaly_key"`
	PostId int `json:"post_id"`
	Content string `json:"content"`
	IsDone bool `json:"is_done"`
}

func FindTask(t *Task) Task {
	var task Task
	db := Init()
	db.Where(t).Find(&task)
	return task
}

func UpdateTask(t *Task) error {
	db := Init()
	rows := db.Model(t).Update(map[string]interface{}{
		"is_done": t.IsDone,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Task (%v) to update", t)
	}
	return nil
}
