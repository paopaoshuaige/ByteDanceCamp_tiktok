package dao

import (
	"gorm.io/gorm"
	"sync"
)

var (
	commentDao  *CommentDAO
	commentOnce sync.Once
)

// 评论信息
// idx_video_id: 查找视频ID对应的所有评论
type Comment struct {
	gorm.Model
	UserID  int64  `gorm:"not null" json:"user_id"`
	VideoID int64  `gorm:"not null; index:idx_video_id" json:"video_id"`
	Content string `gorm:"not null" json:"content"`
}

type CommentDAO struct {
}

func NewCommentDao() *CommentDAO {
	commentOnce.Do(func() {
		commentDao = new(CommentDAO)
	})
	return commentDao
}
