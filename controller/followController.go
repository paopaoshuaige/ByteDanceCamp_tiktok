package controller

import (
	"ByteDanceCamp_tiktok/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RelationListResponse struct {
	Response
	UserList []service.User `json:"user_list,omitempty"`
}

// RelationAction 关注操作
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if actionType == 1 {
		f := service.Follow(toUserId, token)
		switch f {
		case 0:
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "关注成功"})
		case 1:
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "已关注"})
		case 2:
			c.JSON(http.StatusOK, Response{StatusCode: 2, StatusMsg: "关注失败"})
		}
	} else {
		f := service.UnFollow(toUserId, token)
		switch f {
		case 0:
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "取消关注成功"})
		case 1:
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "未关注该用户"})
		case 2:
			c.JSON(http.StatusOK, Response{StatusCode: 2, StatusMsg: "取消关注失败"})
		}
	}
}

// FollowList 关注列表
func FollowList(c *gin.Context) {
	fmt.Println("我执行了")
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	relationListData := service.FollowList(token, userId)
	if relationListData != nil {
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{StatusCode: 0, StatusMsg: "查询成功"},
			UserList: relationListData,
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "查询失败"})
	}
}

// FollowerList 粉丝列表
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	relationListData := service.FollowerList(token, userId)
	if relationListData != nil {
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{StatusCode: 0, StatusMsg: "查询成功"},
			UserList: relationListData,
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "没有粉丝"})
	}
}
