package dao

import (
	"gorm.io/gorm"
	"sync"
)

var (
	followDao  *FollowDAO
	followOnce sync.Once
)

type Follow struct {
	gorm.Model
	FollowerID int64 `gorm:"not null; index:idx_follower" json:"follower_id"` // 被关注
	FollowID   int64 `gorm:"not null; index:idx_follow" json:"follow_id"`     // 关注者
}

type FollowDAO struct {
}

func NewFollowDao() *FollowDAO {
	followOnce.Do(func() {
		followDao = new(FollowDAO)
	})
	return followDao
}

// IsFollow 有没有关注
func (f *FollowDAO) IsFollow(followerId, followId int64) bool {
	return DB.Model(Follow{}).Where("follower_id = ? AND follow_id = ? AND deleted_at IS NULL", followerId, followId).First(&Follow{}).Error == nil
}
