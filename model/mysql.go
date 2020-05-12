package model

import (
	"fmt"

	"github.com/Yukata-team/GoRela-server/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

//DB接続
func Init() *gorm.DB {
	CONNECT := conf.USER + ":" + conf.PASS + "@" + "tcp(" + conf.HOST + ":" + conf.PORT + ")" + "/" + conf.DBNAME + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	db.LogMode(true)
	return db
}
