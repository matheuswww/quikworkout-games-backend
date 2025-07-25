package admin_request

type Noreps struct {
	Time  string `json:"time" binding:"required" validate:"time"`
	NoRep string `json:"norep" binding:"required"`
}

type PutNoreps struct {
	VideoId string   `json:"video_id" binding:"required"`
	EditionId string `json:"edition_id" binding:"required"`
	Category  string `json:"category" validate:"required,category"`
	Sex       string `json:"sex" validate:"required,sex"`
	Noreps  []Noreps `json:"noreps" validate:"dive"`
}
