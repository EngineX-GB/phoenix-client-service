package model

type TelephoneSearchResult struct {
	UserId      string `json:"userId"`
	Username    string `json:"username"`
	Nationality string `json:"nationality"`
	Ethnicity   string `json:"ethnicity"`
	Telephone   string `json:"telephone"`
}

type UsernameSearchResult struct {
	UserId                 string `json:"userId"`
	Username               string `json:"username"`
	CurrentUsername        string `json:"currentUsername"`
	Nationality            string `json:"nationality"`
	Ethnicity              string `json:"ethnicity"`
	UsernameInTodaysRecord bool   `json:"usernameInTodaysRecord"`
}
