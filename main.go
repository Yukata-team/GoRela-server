package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/Yukata-team/GoRela-server/route"

)

func init() {
	//おまじない
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	router := route.Init()

	// サーバー起動
	router.Start(":1323")
}
