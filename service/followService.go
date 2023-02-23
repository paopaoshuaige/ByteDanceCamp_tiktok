package service

import "ByteDanceCamp_tiktok/utils"

// Follow 关注操作
func Follow(toUserId int64, token string) int {
	userId, err := utils.ParseToken(token)
	if err != nil {
		return 2
	}
	return followDao.AddFollow(userId, toUserId)
}

// UnFollow 取关
func UnFollow(toUserId int64, token string) int {
	userId, err := utils.ParseToken(token)
	if err != nil {
		return 2
	}
	return followDao.DeleteFollow(userId, toUserId)
}

// FollowList 关注列表
func FollowList(token string, userId int64) []User {
	userList, status := followDao.QueryFollowById(userId)
	if status != 0 {
		return nil
	}
	var userInfoList []User
	for i := range userList {
		var userInfo *User
		userInfo = QueryUserInfo(token, userList[i])
		userInfoList = append(userInfoList, *userInfo)
	}
	return userInfoList
}

// FollowerList 粉丝列表
func FollowerList(Token string, UserId int64) []User {
	userList := followDao.QueryFollowerById(UserId)
	if userList == nil {
		return nil
	}
	var userInfoList []User
	for i := range userList {
		var userInfo *User
		userInfo = QueryUserInfo(Token, userList[i])
		userInfoList = append(userInfoList, *userInfo)
	}
	return userInfoList
}
