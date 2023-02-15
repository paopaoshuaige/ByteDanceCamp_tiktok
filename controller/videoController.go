package controller

import (
	"ByteDanceCamp_tiktok/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FeedResponse 返回给前端的json结构体
type FeedResponse struct {
	Response
	Data *service.VideoData
}

// Feed 传回视频流
func Feed(c *gin.Context) {
	// 时间戳为该视频现在的时间
	timeStamp := c.Query("latest_time")
	token := c.Query("token")
	feedData := service.NewVideoService(timeStamp, token).Feed()
	if feedData != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: int32(0),
				StatusMsg:  "video feed success",
			},
			Data: feedData,
		})
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(1),
			StatusMsg:  "video feed pull error",
		})
	}
}
