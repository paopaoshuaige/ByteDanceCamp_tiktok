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
	tkGroup.GET("/publish/list/", controller.PublishList)
	// 发布视频
	tkGroup.POST("/publish/action/", controller.Publish)
	// 用户信息
	tkGroup.GET("/user/", controller.QueryUserInfo)
	// 注册
	tkGroup.POST("/user/register/", controller.Register)
	// 登录
	tkGroup.POST("/user/login/", controller.Login)

	// 互动方向
	// 点赞
	tkGroup.POST("/favorite/action/", controller.FavoriteAction)
	// 喜欢列表
	tkGroup.GET("/favorite/list/", controller.FavoriteList)
	// 评论功能
	tkGroup.POST("/comment/action/", controller.CommentAction)
	// 视频评论列表
	tkGroup.GET("/comment/list/", controller.CommentList)

	// 社交接口
	// 关注
	tkGroup.POST("/relation/action/", controller.RelationAction)
	// 用户关注列表
	tkGroup.GET("/relation/follow/list/", controller.FollowList)
	// 用户粉丝列表
	tkGroup.GET("/relation/follower/list/", controller.FollowerList)
}
