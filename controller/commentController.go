package controller

import (
	"ByteDanceCamp_tiktok/service"
	"ByteDanceCamp_tiktok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CommentActionResponse 返回的评论结构体
type CommentActionResponse struct {
	Response
	Id         int64         `json:"id"`
	User       *service.User `json:"user"`
	Content    string        `json:"content"`
	CreateDate string        `json:"create_date"`
}

// CommentListResponse 评论列表
type CommentListResponse struct {
	Response
	CommentList []service.CommentWithAuthor `json:"comment_list,omitempty"`
}

// CommentAction 评论功能
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	commentText := c.Query("comment_text")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	userId, _ := utils.ParseToken(c.Query("token"))
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	commentData := service.CommentAciton(userId, videoId, commentId, actionType, commentText, token)
	if actionType == 1 {
		if commentData != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response:   Response{StatusCode: 0, StatusMsg: "评论成功"},
				Id:         commentData.Id,
				User:       commentData.User,
				Content:    commentData.Content,
				CreateDate: commentData.CreateDate,
			})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "评论失败"})
		}
	} else {
		if commentData != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "删除评论成功"})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "删除评论失败"})
		}
	}
}

// CommentList 评论列表功能
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	commentListData := service.CommentList(token, videoId)
	if commentListData != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0, StatusMsg: "获取评论列表成功"},
			CommentList: commentListData,
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "视频暂无评论"})
	}
}
