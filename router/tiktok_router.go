package router

import (
	"ByteDanceCamp_tiktok/controller"
	"github.com/gin-gonic/gin"
)

func InitTikTokRouter(r *gin.Engine) {
	// 静态路由
	r.Static("static", "./static")

	tkGroup := r.Group("/douyin")
	// 视频推送
	tkGroup.GET("/feed/", controller.Feed)
	// 视频列表
	tkGroup.GET("/publish/list/")
	// 发布视频
	tkGroup.POST("/publish/action/")
	// 用户信息
	tkGroup.GET("/user/", controller.QueryUserInfo)
	// 注册
	tkGroup.POST("/user/register/", controller.Register)
	// 登录
	tkGroup.POST("/user/login/", controller.Login)

	// 互动方向
	// 点赞
	tkGroup.POST("/favorite/action/")
	// 喜欢列表
	tkGroup.GET("/favorite/list/")
	// 评论列表
	tkGroup.POST("/comment/action/")
	// 视频评论列表
	tkGroup.GET("/comment/list/")

	// 社交接口
	// 关注
	tkGroup.POST("/relation/action/")
	// 用户关注列表
	tkGroup.GET("/relatioin/follow/list/")
	// 用户粉丝列表
	tkGroup.GET("/relation/follower/list/")
}
