package participant_request

type GetParticipant struct {
	UserID    					string `form:"userId"`
	EditionId 					string `form:"editionId"`
	CursorCreatedAt   	string `form:"cursorCreatedAt"`
	CursorUserTime  	  string `form:"cursorUserTime"`
	BestTime 						bool   `form:"bestTime"`
	WorstTime 					bool   `form:"worstTime"`
}