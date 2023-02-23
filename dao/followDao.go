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

// AddFollow 新增关注
func (f *FollowDAO) AddFollow(followerId, followId int64) int {
	follow := Follow{
		FollowerID: followerId,
		FollowID:   followId,
	}
	// 已经关注了
	if f.IsFollow(followerId, followId) {
		return 1
	}
	if err := DB.Create(&follow).Error; err != nil {
		return 2
	}
	if DB.Model(&User{}).Where("id = ?", followerId).Update("follow_count", gorm.Expr("follow_count + 1")).Error != nil {
		return 2
	}
	if DB.Model(&User{}).Where("id = ?", followId).Update("follower_count", gorm.Expr("follower_count + 1")).Error != nil {
		return 2
	}
	return 0
}

// DeleteFollow 取消关注
func (f *FollowDAO) DeleteFollow(followerId, followId int64) int {
	// 已经取关了
	if !f.IsFollow(followerId, followId) {
		return 1
	}
	if err := DB.Model(&Follow{}).Where("follower_id = ? AND follow_id = ? AND deleted_at IS NULL", followerId, followId).Delete(&Follow{}).Error; err != nil {
		return 2
	}
	if DB.Model(&User{}).Where("id = ?", followerId).Update("follow_count", gorm.Expr("follow_count - 1")).Error != nil {
		return 2
	}
	if DB.Model(&User{}).Where("id = ?", followId).Update("follower_count", gorm.Expr("follower_count - 1")).Error != nil {
		return 2
	}
	return 0
}

// QueryFollowById 根据id查询所有关注
func (f *FollowDAO) QueryFollowById(userId int64) ([]int64, int) {
	var userList []int64
	if err := DB.Model(&Follow{}).Select("follow_id").Where("follower_id = ? AND deleted_at IS NULL", userId).Find(&userList).Error; err != nil {
		return nil, 1
	}
	return userList, 0
}

// QueryFollowerById 根据id查询所有的粉丝
func (f *FollowDAO) QueryFollowerById(userId int64) []int64 {
	var userList []int64
	if err := DB.Model(&Follow{}).Select("follower_id").Where("follow_id = ? AND deleted_at IS NULL", userId).Find(&userList).Error; err != nil {
		return nil
	}
	return userList
}

// IsFollow 有没有关注
func (f *FollowDAO) IsFollow(followerId, followId int64) bool {
	return DB.Model(Follow{}).Where("follower_id = ? AND follow_id = ? AND deleted_at IS NULL", followerId, followId).First(&Follow{}).Error == nil
}
