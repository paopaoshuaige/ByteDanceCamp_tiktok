package service

import (
	"ByteDanceCamp_tiktok/utils"
	"strconv"
)

// Favorite 点赞和取消点赞
func Favorite(token, videoId string, actionType int64) int {
	Id, err := utils.ParseToken(token)
	if err != nil {
		return -1
	}
	VideoId, _ := strconv.ParseInt(videoId, 10, 64)
	if actionType == 1 {
		return favoriteDao.AddFavorite(Id, VideoId)
	} else {
		return favoriteDao.DeleteFavorite(Id, VideoId)
	}
	return -1
}

// FavoriteList 点赞列表
func FavoriteList(userId, token string) []VideoDisplay {
	User, err := utils.ParseToken(token)
	if err != nil {
		return nil
	}
	queryUser, _ := strconv.ParseInt(userId, 10, 64)
	// 视频id列表
	videoIdList, status := favoriteDao.QueryFavoriteById(queryUser)
	if status != 0 {
		return nil
	}
	videoDisplayList := make([]VideoDisplay, 0, 30)
	for i := range videoIdList {
		var videoDisplay VideoDisplay
		video, err := videoDao.QueryVideoById(videoIdList[i])
		if err != nil {
			return nil
		}
		videoDisplay = VideoDisplay{
			Id:            video.ID,
			Title:         video.Title,
			CreatedAt:     video.CreatedAt,
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
		}
		videoDisplay.Author = QueryUserInfo(token, video.AuthorId)
		videoDisplay.IsFavorite = favoriteDao.IsFavorite(User, videoIdList[i])
		videoDisplayList = append(videoDisplayList, videoDisplay)
	}
	return videoDisplayList
}
