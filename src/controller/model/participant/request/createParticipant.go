package participant_request

type CreateParticipant struct {
	Title 		string 				 `json:"title" binding:"required"`
	Size      int64					 `json:"size" binding:"required"`
	UserTime  string         `json:"userTime" binding:"required" validate:"time"`
	Sex     	string 				 `json:"sex" binding:"required" validate:"required,sex"`
}