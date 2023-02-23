package dao

import (
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	FavoriteDao  *FavoriteDAO
	FavoriteOnce sync.Once
)

type Favorite struct {
	gorm.Model
	UserID  int64 `gorm:"not null; index:idx_userid_videoid" json:"user_id"`
	VideoID int64 `gorm:"not null; index:idx_userid_videoid" json:"video_id"`
}

type FavoriteDAO struct {
}

func NewFavoriteDao() *FavoriteDAO {
	FavoriteOnce.Do(func() {
		FavoriteDao = new(FavoriteDAO)
	})
	return FavoriteDao
}

// AddFavorite 新增点赞
func (f *FavoriteDAO) AddFavorite(userId, videoId int64) int {
	if f.IsFavorite(userId, videoId) {
		return 1
	}
	if err := DB.Create(&Favorite{UserID: userId, VideoID: videoId}).Error; err != nil {
		return -1
	}
	if DB.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + 1")).Error != nil {
		return -1
	}
	return 0
}

// DeleteFavorite 是否取消点赞
func (f *FavoriteDAO) DeleteFavorite(userId, videoId int64) int {
	if !f.IsFavorite(userId, videoId) {
		return 1
	}
	if err := DB.Where("user_id = ? AND video_id = ? AND deleted_at IS NULL", userId, videoId).Delete(&Favorite{}).Error; err != nil {
		return -1
	}
	if DB.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - 1")).Error != nil {
		return -1
	}
	return 0
}

// QueryFavoriteById 查询点赞列表（id）
func (f *FavoriteDAO) QueryFavoriteById(userId int64) ([]int64, int) {
	var videoIdList []int64
	if err := DB.Model(&Favorite{}).Select("video_id").Where("user_id = ? AND deleted_at IS NULL", userId).Find(&videoIdList).Error; err != nil {
		log.Fatal("获取视频列表失败：", err)
		return nil, 1
	}
	return videoIdList, 0
}

// IsFavorite 查询是否点赞
func (f *FavoriteDAO) IsFavorite(userId, videoId int64) bool {
	return DB.Model(&Favorite{}).Where("user_id = ? AND video_id = ? AND deleted_at IS NULL", userId, videoId).First(&Favorite{}).Error == nil
}
