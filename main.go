package main

import (
	"github.com/Yukata-team/GoRela-server/route"
)

func main() {
	router := route.Init()
	// サーバー起動
	router.Start(":1323")
}
