package service

import (
	"ByteDanceCamp_tiktok/dao"
	"ByteDanceCamp_tiktok/utils"
	"mime/multipart"
	"path/filepath"
	"strconv"
)

type PublishService struct {
	Token      string
	Title      string
	Fileheader *multipart.FileHeader
}

type PublishListService struct {
	UserId string
	Token  string
}

type PublishListData struct {
	VideoList []VideoDisplay `json:"video_list,omitempty"`
}

func NewPublishService(token, title string, fileheader *multipart.FileHeader) *PublishService {
	return &PublishService{Token: token, Title: title, Fileheader: fileheader}
}

func NewPublishListService(userid, token string) *PublishListService {
	return &PublishListService{UserId: userid, Token: token}
}

func (p *PublishService) Publish() error {

	userId, err := utils.ParseToken(p.Token)
	if err != nil {
		return err
	}
	// 封面，有点bug，使用默认抖音的
	videoName := filepath.Base(p.Fileheader.Filename)
	// videoPath := filepath.Join("./static/video", videoName)
	// coverPath := "./static/covers/" + videoName
	// coverName, _ := utils.CoverGenerator(videoPath, coverPath)

	playURL := "http://172.20.10.3:8989/static/video/" + videoName
	coverUrl := "http://172.20.10.3:8989/static/covers/img.png"
	video := &dao.Video{
		AuthorId: userId,
		Title:    p.Title,
		PlayURL:  playURL,
		CoverURL: coverUrl,
	}
	if err := videoDao.AddVideo(video); err != nil {
		return err
	}
	return nil
}

// PublishList 根据作者id和传入的id查询视频记录，并倒序列出
func (p *PublishListService) PublishList() ([]VideoDisplay, error) {
	userId, _ := strconv.ParseInt(p.UserId, 10, 64)
	videoList, err := videoDao.QueryVideoByUserId(userId)
	if err != nil {
		return nil, err
	}
	//获得作者信息
	var userInfo *User
	videoDisplayList := make([]VideoDisplay, 0, 30)
	userInfo = NewUserInfoService(p.Token, userId).QueryUserInfo()
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
