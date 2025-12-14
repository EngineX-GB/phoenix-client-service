package model

type LinkRequest struct {
	UserId1   string `json:"userId1"`
	UserId2   string `json:"userId2"`
	InputType string `json:"inputType"`
	Notes     string `json:"notes"`
}
