package controller

import (
	"ByteDanceCamp_tiktok/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

// PublishResponse 当前用户视频发布列表json响应
type PublishResponse struct {
	Response
	Data *service.PublishListData
}

// Publish 上传视频
func Publish(c *gin.Context) {
	title := c.PostForm("title")
	token := c.PostForm("token")
	data, err := c.FormFile("data")
	if token == "" || title == "" || data == nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "有必须字段为空"})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	if err = service.NewPublishService(token, title, data).Publish(); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 获取文件名字加后缀
	fileSuffix := filepath.Base(data.Filename)
	saveFile := filepath.Join("./static/video", fileSuffix)
	// 存放到本地
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "uploaded success",
	})
}

// PublishList 获取用户视频列表
func PublishList(c *gin.Context) {
	//获取用户id
	userId := c.Query("user_id")
	token := c.Query("token")
	publishListData, err := service.NewPublishListService(userId, token).PublishList()
	fmt.Println(publishListData)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "获取失败",
		})
	}
	c.JSON(http.StatusOK, PublishResponse{
		Response: Response{StatusCode: 0},
		Data:     publishListData,
	})
}
