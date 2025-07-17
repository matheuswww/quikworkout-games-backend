package participant_request

type CreateParticipant struct {
	Title 		string 				 `json:"title" binding:"required"`
	Size      int64					 `json:"size" binding:"required"`
	Sex     	string 				 `json:"sex" binding:"required" validate:"required,sex"`
}