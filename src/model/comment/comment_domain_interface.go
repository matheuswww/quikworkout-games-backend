package comment_domain

type CommentDomainInterface interface {
	GetCommentId() string
	SetCommentId(string)
	GetVideoId() string
	SetVideoId(string)
	GetParentId() string
	SetParentId(string)
	GetAnswerId() string
	SetAnswerId(string)
	GetUserId() string
	SetUserId(string)
	GetVideoComment() string
	SetVideoComment(string)
	GetCreatedAt() string
	SetCreatedAt(string)
}

func NewCommentDomain(
	commentId, videoId, parentId, answerId, userId, videoComment,
	createdAt string,
) CommentDomainInterface {
	return &comment{
		commentId:    commentId,
		videoId:      videoId,
		parentId:     parentId,
		answerId: 		answerId,
		userId:       userId,
		videoComment: videoComment,
		createdAt:    createdAt,
	}
}
