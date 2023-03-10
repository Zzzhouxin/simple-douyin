package main

import (
	"github.com/RaymondCode/simple-douyin/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	//go service.RunMessageServer()

	dao.Init()
	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
