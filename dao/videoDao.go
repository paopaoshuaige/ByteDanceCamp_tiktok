package dao

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	videoDao  *VideoDAO
	videoOnce sync.Once
)

// Video 数据表Video结构
type Video struct {
	gorm.Model
	AuthorId      int64  `json:"author_id" gorm:"not null; index:idx_author_id"` // 视频作者信息
	CommentCount  int64  `json:"comment_count"`                                  // 视频的评论总数
	CoverURL      string `json:"cover_url"`                                      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"`                                 // 视频的点赞总数
	ID            int64  `json:"id"`                                             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`                                    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`                                       // 视频播放地址
	Title         string `json:"title"`                                          // 视频标题
}

type VideoDAO struct {
}

func NewVideoDao() *VideoDAO {
	videoOnce.Do(func() {
		videoDao = new(VideoDAO)
	})
	return videoDao
}

// AddVideo 添加视频
func (v *VideoDAO) AddVideo(video *Video) error {
	return DB.Create(video).Error
}

// QueryVideoById 根据id搜索视频
func (vd *VideoDAO) QueryVideosById(id int64) (*Video, error) {
	var video Video
	err := DB.Model(&Video{}).Where("id = ?", id).Find(&video).Error
	return &video, err
}

func (vd *VideoDAO) QueryVideosByUserId(id int64) ([]Video, error) {
	videoList := make([]Video, 0, 30)
	err := DB.Model(&Video{}).Where("author_id = ?", id).Order("created_at desc").Limit(30).Find(&videoList).Error
	return videoList, err
}

// QueryVideoByLatestTime 查询视频最新时间
func (vd *VideoDAO) QueryVideoByLatestTime(latestTime int64) []Video {
	var videoList []Video
	timeStamp := time.UnixMilli(latestTime)
	// 创建时间倒序
	DB.Model(&Video{}).Where("created_at>=?", timeStamp).Order("created_at desc").Limit(30).Find(&videoList)

	// 没有视频
	if len(videoList) == 0 {
		return nil
	}

	return videoList
}
