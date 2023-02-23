package service

import "ByteDanceCamp_tiktok/dao"

type CommentActionService struct {
	UserId      int64
	Token       string
	VideoId     int64
	ActionType  int64
	CommentText string
	CommentId   int64
}

type CommentListService struct {
	Token   string
	VideoId int64
}

type CommentWithAuthor struct {
	Id         int64  `json:"id"`
	User       *User  `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type CommentActionData struct {
	CommentWithAuthor `json:"comment,omitempty"`
}

func CommentAciton(userId, videoId, commentId, actionType int64, commentText, token string) *CommentActionData {
	var commentData *CommentActionData
	if actionType == 1 {
		var userInfo *User
		var comment *dao.Comment
		comment = commentDao.AddComment(userId, videoId, commentText)
		if comment == nil {
			return nil
		}
		userInfo = QueryUserInfo(token, userId)
		commentData = &CommentActionData{CommentWithAuthor: CommentWithAuthor{
			Id:         int64(comment.ID),
			User:       userInfo,
			Content:    commentText,
			CreateDate: comment.CreatedAt.Format("2006-01-02 15:04:05")[5:10],
		}}
	} else {
		err := commentDao.DeleteComment(commentId)
		if err != nil {
			return nil
		}
	}

	return commentData
}

func CommentList(Token string, VideoId int64) []CommentWithAuthor {
	commentList := commentDao.QueryCommentListByVideoId(VideoId)
	if len(commentList) == 0 {
		return nil
	}
	var commentListData []CommentWithAuthor
	for i := range commentList {
		var commentWithAuthor CommentWithAuthor
		var userInfo *User
		userInfo = QueryUserInfo(Token, commentList[i].UserID)
		commentWithAuthor = CommentWithAuthor{
			Id:         int64(commentList[i].ID),
			User:       userInfo,
			Content:    commentList[i].Content,
			CreateDate: commentList[i].CreatedAt.Format("2006-01-02 15:04:05")[5:10],
		}
		commentListData = append(commentListData, commentWithAuthor)
	}
	return commentListData
}
