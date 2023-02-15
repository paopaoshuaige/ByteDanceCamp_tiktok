package service

import (
	"ByteDanceCamp_tiktok/dao"
	"ByteDanceCamp_tiktok/utils"
	"strconv"
	"time"
)

// VideoService 视频验证
type VideoService struct {
	TimeStamp string // 时间戳
	Token     string // token
}

// VideoDisplay 视频信息
type VideoDisplay struct {
	Id            int64     `json:"id,omitempty" gorm:"primaryKey"`     // 视频id
	Author        *User     `json:"author" gorm:"-"`                    // 作者信息
	PlayUrl       string    `json:"play_url" json:"play_url,omitempty"` // 视频播放地址
	CoverUrl      string    `json:"cover_url,omitempty"`                // 视频封面地址
	FavoriteCount int64     `json:"favorite_count,omitempty"`           // 视频的点赞总数
	CommentCount  int64     `json:"comment_count,omitempty"`            // 视频的评论总数
	IsFavorite    bool      `json:"is_favorite,omitempty"`              // 是否喜欢
	Title         string    `json:"title"`                              // 标题
	CreatedAt     time.Time `json:"-"`                                  // 创建时间
	UpdatedAt     time.Time `json:"-"`                                  // 更新时间
}

var (
	videoDao  = dao.NewVideoDao()
	likeDao   = dao.NewLikeDao()
	userDao   = dao.NewUserInfoDAO()
	followDao = dao.NewFollowDao()
)

// VideoData 视频信息
type VideoData struct {
	NextTime  int64          `json:"next_time"`            // 这次视频最近投稿时间
	VideoList []VideoDisplay `json:"video_list,omitempty"` // 视频列表
}

// NewVideoService 创建一个新的videosvc
func NewVideoService(timeStamp, token string) *VideoService {
	return &VideoService{TimeStamp: timeStamp, Token: token}
}

// Feed 传回三十个视频流
func (v *VideoService) Feed() *VideoData {
	IntTime, _ := strconv.ParseInt(v.TimeStamp, 10, 64)
	userToken, _ := utils.ParseToken(v.Token)
	// 通过创建时间倒序返回三十个视频
	videoList := videoDao.QueryVideoByLatestTime(IntTime)
	// 视频作者信息
	videoDisplayList := make([]VideoDisplay, 0, 30)
	for video := range videoList {
		videoDisplay := VideoDisplay{
			Id:            videoList[video].ID,
			Title:         videoList[video].Title,
			CreatedAt:     videoList[video].CreatedAt,
			PlayUrl:       videoList[video].PlayURL,
			CoverUrl:      videoList[video].CoverURL,
			FavoriteCount: videoList[video].FavoriteCount,
			CommentCount:  videoList[video].CommentCount,
		}
		// 视频作者信息
		videoDisplay.Author = NewUserInfoService(v.Token, videoList[video].AuthorId).QueryUserInfo()
		// 目前用户是否喜欢了该视频
		videoDisplay.IsFavorite = likeDao.IsFavorite(userToken, videoDisplay.Id)
		videoDisplayList = append(videoDisplayList, videoDisplay)
	}
	// 返回这次视频最近的投稿时间 - 1，下次即可获取比这次视频旧的视频
	nextTime := videoList[len(videoList)-1].CreatedAt.UnixMilli() - 1
	// nextTime := videoList[len(videoList)-1].CreatedAt.UnixMilli()
	return &VideoData{NextTime: nextTime, VideoList: videoDisplayList}
}
