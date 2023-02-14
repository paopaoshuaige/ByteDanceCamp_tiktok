package main

import (
	"ByteDanceCamp_tiktok/dao"
	"ByteDanceCamp_tiktok/router"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.InitDB()
	r := gin.Default()
	router.InitTikTokRouter(r)
	r.Run(":8989")
}
