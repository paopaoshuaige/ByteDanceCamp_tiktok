package controller

import (
	"ByteDanceCamp_tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteResponse 是否喜欢
type FavoriteResponse struct {
	Response
}

// FavoriteListResponse 喜欢列表
type FavoriteListResponse struct {
	Response
	VideoList []service.VideoDisplay `json:"video_list,omitempty"`
}

// FavoriteAction 点赞操作
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if actionType == 1 {
		if f := service.Favorite(token, videoId, actionType); f == 0 {
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "点赞成功"})
		} else if f == 1 {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "已经点过赞了"})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 2, StatusMsg: "点赞失败"})
		}
	} else {
		if f := service.Favorite(token, videoId, actionType); f == 0 {
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "已经取消点赞"})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "取消点赞失败"})
		}
	}
}

// FavoriteList 点赞列表
func FavoriteList(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	favoriteListData := service.FavoriteList(userId, token)
	if favoriteListData != nil {
		c.JSON(http.StatusOK, FavoriteListResponse{
			Response:  Response{StatusCode: 0, StatusMsg: "查询成功"},
			VideoList: favoriteListData,
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "查询失败"})
	}
}
