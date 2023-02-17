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
	ID            int64  `json:"id,omitempty"`             // 用户id
	IsFollow      bool   `json:"is_follow,omitempty"`      // true-已关注，false-未关注
	Name          string `json:"name,omitempty"`           // 用户名称
}

// UserRegisterService 用户注册
type UserRegisterService struct {
	Username string
	Password string
}

// UserLoginService 用户登录
type UserLoginService struct {
	Username string
	Password string
}

// UserRegisterData 用户注册信息
type UserRegisterData struct {
	Id    int64  `json:"user_id"`
	Token string `json:"token"`
}

// UserLoginData 用户登录信息
type UserLoginData struct {
	Id    int64  `json:"user_id"`
	Token string `json:"token"`
}

// UserInfoService 视频作者信息
type UserInfoService struct {
	Token string
	Id    int64
}

// NewUserInfoService 返回一个视频作者信息
func NewUserInfoService(token string, id int64) *UserInfoService {
	return &UserInfoService{Token: token, Id: id}
}

// NewRegisterService 返回一个用户注册信息
func NewRegisterService(username, pass string) *UserRegisterService {
	return &UserRegisterService{Username: username, Password: pass}
}

// NewLoginService 返回用户登录信息
func NewLoginService(username, pass string) *UserLoginService {
	return &UserLoginService{Username: username, Password: pass}
}

func (u *UserRegisterService) Register() (*UserRegisterData, int) {
	user := dao.User{}
	salt := make([]byte, 32)
	for i := range salt { // 生成随机数
		salt[i] = byte(rand.Intn(256))
	}
	user.Salt = salt
	// 对用户密码加密
	user.Password = utils.Encrypt(u.Password, salt)
	user.Name = u.Username
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
func (u *UserLoginService) Login() (*UserLoginData, int) {
	// 查询用户名是否重复
	user, f := userDao.QueryUserByName(u.Username)
	if !f {
		return nil, 1
	}
	// 加密
	password := utils.Encrypt(u.Password, user.Salt)
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
func (u *UserInfoService) QueryUserInfo() *User {
	// 解析token
	userToken, err := utils.ParseToken(u.Token)
	if err != nil {
		return nil
	}
	// 根据id查询信息
	user, err := userDao.QueryUserById(u.Id)
	if err != nil {
		return nil
	}
	userInfo := &User{
		ID:            user.ID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
	}
	// 查询刷视频的用户是否关注了该视频作者
	userInfo.IsFollow = followDao.IsFollow(userToken, u.Id)
	return userInfo
}
