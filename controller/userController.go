package controller

import (
	"ByteDanceCamp_tiktok/service"
	"ByteDanceCamp_tiktok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserRegisterResponse 用户注册返回的json
type UserRegisterResponse struct {
	Response
	Id    int64  `json:"user_id,omitempty"`
	Token string `json:"token"`
}

// UserLoginResponse 用户登录返回的json
type UserLoginResponse struct {
	Response
	Id    int64  `json:"user_id,omitempty"`
	Token string `json:"token"`
}

// UserInfoResponse 用户查询返回的响应
type UserInfoResponse struct {
	Response
	User *service.User `json:"user"`
}

// Register 注册
func Register(c *gin.Context) {
	username := c.Query("username")
	pass := c.Query("password")

	if !utils.CheckName(username) {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "username Does not meet specifications",
		})
		return
	}

	if !utils.CheckPass(pass) {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "password Does not meet specifications",
		})
	}

	userRegisterData, flag := service.Register(username, pass)

	if flag == 0 {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "Register success",
			},
			Id:    userRegisterData.Id,
			Token: userRegisterData.Token,
		})
	} else if flag == 1 {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "Duplicate username",
		})
	} else if flag == 2 {
		c.JSON(http.StatusOK, Response{
			StatusCode: 2,
			StatusMsg:  "Register error",
		})
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 3,
			StatusMsg:  "token error",
		})
	}
}

// Login 登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if !utils.CheckName(username) {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "username Does not meet specifications",
		})
		return
	}

	if !utils.CheckPass(password) {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "password Does not meet specifications",
		})
	}

	userLoginData, flag := service.Login(username, password)

	if flag == 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "Login success",
			},
			Id:    userLoginData.Id,
			Token: userLoginData.Token,
		})
	} else if flag == 1 {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "username is incorrect",
		})
	} else if flag == 2 {
		c.JSON(http.StatusOK, Response{
			StatusCode: 2,
			StatusMsg:  "password is incorrect",
		})
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 3,
			StatusMsg:  "token error",
		})
	}
}

// QueryUserInfo 查询信息
func QueryUserInfo(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	user := service.QueryUserInfo(token, userId)
	if user == nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "Query failed",
		})
	} else {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "Query success",
			},
			User: user,
		})
	}
}
