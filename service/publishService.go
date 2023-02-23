package service

import (
	"ByteDanceCamp_tiktok/dao"
	"ByteDanceCamp_tiktok/utils"
	"mime/multipart"
	"path/filepath"
	"strconv"
)

func Publish(Token, Title string, Fileheader *multipart.FileHeader) error {

	userId, err := utils.ParseToken(Token)
	if err != nil {
		return err
	}
	// 封面，有点bug，使用默认抖音的
	videoName := filepath.Base(Fileheader.Filename)
	// videoPath := filepath.Join("./static/video", videoName)
	// coverPath := "./static/covers/" + videoName
	// coverName, _ := utils.CoverGenerator(videoPath, coverPath)

	playURL := "http://" + PlayURL + ":8989/static/video/" + videoName
	coverUrl := "http://" + PlayURL + ":8989/static/covers/img.png"
	video := &dao.Video{
		AuthorId: userId,
		Title:    Title,
		PlayURL:  playURL,
		CoverURL: coverUrl,
	}
	if err := videoDao.AddVideo(video); err != nil {
		return err
	}
	return nil
}

// PublishList 根据作者id和传入的id查询视频记录，并倒序列出
func PublishList(UserId, Token string) ([]VideoDisplay, error) {
	userId, _ := strconv.ParseInt(UserId, 10, 64)
	videoList, err := videoDao.QueryVideoByUserId(userId)
	if err != nil {
		return nil, err
	}
	//获得作者信息
	var userInfo *User
	videoDisplayList := make([]VideoDisplay, 0, 30)
	userInfo = QueryUserInfo(Token, userId)
	for i := range videoList {
		videoDisplay := VideoDisplay{
			Id:            videoList[i].ID,
			Title:         videoList[i].Title,
			Author:        userInfo,
			CreatedAt:     videoList[i].CreatedAt,
			PlayUrl:       videoList[i].PlayURL,
			CoverUrl:      videoList[i].CoverURL,
			FavoriteCount: videoList[i].FavoriteCount,
			CommentCount:  videoList[i].CommentCount,
		}
		videoDisplayList = append(videoDisplayList, videoDisplay)
	}
	return videoDisplayList, err
}
