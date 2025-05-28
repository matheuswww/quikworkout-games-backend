package comment_domain

type comment struct {
	commentId    string
	videoId      string
	parentId     string
	answerId       string
	userId       string
	videoComment string
	createdAt    string
}

func (c *comment) GetCommentId() string {
	return c.commentId
}

func (c *comment) SetCommentId(id string) {
	c.commentId = id
}

func (c *comment) GetVideoId() string {
	return c.videoId
}

func (c *comment) SetVideoId(id string) {
	c.videoId = id
}

func (c *comment) GetParentId() string {
	return c.parentId
}

func (c *comment) SetParentId(id string) {
	c.parentId = id
}

func (c *comment) GetAnswerId() string {
	return c.answerId
}

func (c *comment) SetAnswerId(id string) {
	c.answerId = id
}

func (c *comment) GetUserId() string {
	return c.userId
}

func (c *comment) SetUserId(id string) {
	c.userId = id
}

func (c *comment) GetVideoComment() string {
	return c.videoComment
}

func (c *comment) SetVideoComment(comment string) {
	c.videoComment = comment
}

func (c *comment) GetCreatedAt() string {
	return c.createdAt
}

func (c *comment) SetCreatedAt(t string) {
	c.createdAt = t
}
