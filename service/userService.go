package service

import (
	"ByteDanceCamp_tiktok/dao"
	"ByteDanceCamp_tiktok/utils"
	"math/rand"
	"reflect"
)

// User 返回的用户信息表
type User struct {
	FollowCount   int64  `json:"follow_count,omitempty"`   // 关注总数
	FollowerCount int64  `json:"follower_coun,omitemptyt"` // 粉丝总数
	Id            int64  `json:"id,omitempty"`             // 用户id
	IsFollow      bool   `json:"is_follow,omitempty"`      // true-已关注，false-未关注
	Name          string `json:"name,omitempty"`           // 用户名称
}

// UserRegisterData 用户注册信息
type UserRegisterData struct {
	Id    int64  `json:"user_id,omitempty"`
	Token string `json:"token"`
}

// UserLoginData 用户登录信息
type UserLoginData struct {
	Id    int64  `json:"user_id,omitempty"`
	Token string `json:"token"`
}

// Register 注册
func Register(Username, Password string) (*UserRegisterData, int) {
	user := dao.User{}
	salt := make([]byte, 32)
	for i := range salt { // 生成随机数
		salt[i] = byte(rand.Intn(256))
	}
	user.Salt = salt
	// 对用户密码加密
	user.Password = utils.Encrypt(Password, salt)
	user.Name = Username
	flag := userDao.AddUser(&user)
	if flag != 0 {
		return nil, flag
	}
	token, err := utils.SignToken(user.ID, user.Name)
	if err == nil {
		return &UserRegisterData{Id: user.ID, Token: token}, flag
	}
	return nil, 3
}

// Login 验证是否登陆成功
func Login(Username, Password string) (*UserLoginData, int) {
	// 查询用户名是否重复
	user, f := userDao.QueryUserByName(Username)
	if !f {
		return nil, 1
	}
	// 加密
	password := utils.Encrypt(Password, user.Salt)
	// 判断密码是否正确
	if !reflect.DeepEqual(password, user.Password) {
		return nil, 2
	}
	token, err := utils.SignToken(user.ID, user.Name)
	if err != nil {
		return nil, 3
	}
	return &UserLoginData{Id: user.ID, Token: token}, 0
}

// QueryUserInfo 查询用户信息
func QueryUserInfo(token string, userList int64) *User {
	// 解析token
	userToken, err := utils.ParseToken(token)
	if err != nil {
		return nil
	}
	// 根据id查询信息
	user, err := userDao.QueryUserById(userList)
	if err != nil {
		return nil
	}
	userInfo := &User{
		Id:            user.ID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
	}
	// 查询刷视频的用户是否关注了该视频作者
	userInfo.IsFollow = followDao.IsFollow(userToken, userList)
	return userInfo
}
