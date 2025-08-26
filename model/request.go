package model

type SearchRequest struct {
	Username       string `json:"username"`
	UserId         string `json:"userId"`
	Nationality    string `json:"nationality"`
	Region         string `json:"region"`
	AvailableToday bool   `json:"availableToday"`
}
