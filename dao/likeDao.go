package dao

import (
	"gorm.io/gorm"
	"sync"
)

var (
	likeDao  *LikeDAO
	likeOnce sync.Once
)

// LikeDAO 用户点赞表
type LikeDAO struct {
}

// Like 点赞表
type Like struct {
	gorm.Model
	UserID  int64 `gorm:"not null; index:idx_userid_videoid" json:"user_id"`
	VideoID int64 `gorm:"not null; index:idx_userid_videoid" json:"video_id"`
}

func NewLikeDao() *LikeDAO {
	likeOnce.Do(func() {
		likeDao = new(LikeDAO)
	})
	return likeDao
}

// IsFavorite 是否点赞
func (f *LikeDAO) IsFavorite(userId, videoId int64) bool {
	return DB.Model(&Like{}).Where("user_id = ? AND video_id = ? AND deleted_at IS NULL", userId, videoId).First(&Like{}).Error == nil
}
