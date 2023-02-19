package controller

import (
	"ByteDanceCamp_tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FeedResponse 返回给前端的json结构体
type FeedResponse struct {
	Response
	NextTime  int64                  `json:"next_time"`            // 这次视频最近投稿时间
	VideoList []service.VideoDisplay `json:"video_list,omitempty"` // 视频列表
}

// Feed 传回视频流
func Feed(c *gin.Context) {
	// 时间戳为该视频现在的时间
	timeStamp := c.Query("latest_time")
	token := c.Query("token")
	NextTime, VideoList := service.Feed(timeStamp, token)
	if VideoList != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: int32(0),
				StatusMsg:  "video feed success",
			},
			NextTime:  NextTime,
			VideoList: VideoList,
		})
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(1),
			StatusMsg:  "video feed pull error",
		})
	}
}
