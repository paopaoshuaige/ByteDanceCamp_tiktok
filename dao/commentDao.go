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

func (c *CommentDAO) AddComment(userId, videoId int64, content string) *Comment {
	comment := &Comment{
		UserID:  userId,
		VideoID: videoId,
		Content: content,
	}
	if err := DB.Create(comment).Error; err != nil {
		return nil
	}
	if DB.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + 1")).Error != nil {
		return nil
	}
	return comment
}

func (c *CommentDAO) DeleteComment(commentId int64) error {
	var comment Comment
	if err := DB.Model(&Comment{}).Where("id = ? AND deleted_at IS NULL", commentId).First(&comment).Error; err != nil {
		return err
	}
	if err := DB.Model(&Video{}).Where("id = ?", comment.VideoID).Update("comment_count", gorm.Expr("comment_count - 1")).Error; err != nil {
		return err
	}
	if err := DB.Model(&Comment{}).Where("id = ? AND deleted_at IS NULL", commentId).Delete(&Comment{}).Error; err != nil {
		return err
	}
	return nil
}

// QueryCommentListByVideoId 查询该视频的评论列表
func (c *CommentDAO) QueryCommentListByVideoId(videoId int64) []Comment {
	var commentList []Comment
	if err := DB.Model(&Comment{}).Where("video_id = ? AND deleted_at IS NULL", videoId).Order("created_at ASC").Find(&commentList).Error; err != nil {
		return nil
	}
	return commentList
}
