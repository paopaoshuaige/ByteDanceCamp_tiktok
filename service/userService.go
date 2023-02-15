package service

import "ByteDanceCamp_tiktok/utils"

// User
type User struct {
	FollowCount   int64  `json:"follow_count,omitempty"`   // 关注总数
	FollowerCount int64  `json:"follower_coun,omitemptyt"` // 粉丝总数
	ID            int64  `json:"id,omitempty"`             // 用户id
	IsFollow      bool   `json:"is_follow,omitempty"`      // true-已关注，false-未关注
	Name          string `json:"name,omitempty"`           // 用户名称
}

// UserInfoService 用户信息
type UserInfoService struct {
	token string
	Id    int64
}

func NewUserInfoService(token string, id int64) *UserInfoService {
	return &UserInfoService{token: token, Id: id}
}

// QueryUserInfo 查询用户信息
func (u *UserInfoService) QueryUserInfo() *User {
	// 解析token(看视频的人的)
	userToken, err := utils.ParseToken(u.token)
	if err != nil {
		return nil
	}
	// 根据id查询信息(视频作者的)
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
