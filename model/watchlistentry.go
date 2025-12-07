package model

type WatchListEntry struct {
	UserId      string `json:"userId"`
	Username    string `json:"username"`
	Nationality string `json:"nationality"`
	Telephone   string `json:"telephone"`
	Location    string `json:"location"`
	Rate        int    `json:"rate"`
}
