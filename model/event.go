package model

type RoomEvent struct {
	Revealed bool  `json:"revealed"`
	ResetTs  int64 `json:"resetTs"`
}
