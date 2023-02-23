package service

import "ByteDanceCamp_tiktok/dao"

var (
	videoDao    = dao.NewVideoDao()
	likeDao     = dao.NewLikeDao()
	userDao     = dao.NewUserInfoDAO()
	followDao   = dao.NewFollowDao()
	favoriteDao = dao.NewFavoriteDao()
	commentDao  = dao.NewCommentDao()
)
