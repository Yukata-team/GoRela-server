package main

import (
	"./route"
	"github.com/Sirupsen/logrus"
	// "github.com/labstack/echo/engine/fasthttp"
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
