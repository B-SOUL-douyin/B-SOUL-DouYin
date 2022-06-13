package main

import (
	"github.com/B-SOUL-douyin/B-SOUL-DouYin/video/dal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
