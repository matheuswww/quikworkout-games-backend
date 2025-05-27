package comment_domain

type CommentDomainInterface interface {
	GetCommentId() string
	SetCommentId(string)
	GetVideoId() string
	SetVideoId(string)
	GetParentId() string
	SetParentId(string)
	GetUserId() string
	SetUserId(string)
	GetVideoComment() string
	SetVideoComment(string)
	GetCreatedAt() string
	SetCreatedAt(string)
}

func NewCommentDomain(
	commentId, videoId, parentId, userId, videoComment,
	createdAt string,
) CommentDomainInterface {
	return &comment{
		commentId:    commentId,
		videoId:      videoId,
		parentId:     parentId,
		userId:       userId,
		videoComment: videoComment,
		createdAt:    createdAt,
	}
}
