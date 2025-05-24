package participant_request

type CreateParticipant struct {
	Title 		string 				 `json:"title" binding:"required"`
	Instagram string 				 `json:"instagram" binding:"required"`
	Size      int64						`json:"size" binding:"required"`
}