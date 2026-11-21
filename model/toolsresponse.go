package model

type TelephoneSearchResult struct {
	UserId      string `json:"userId"`
	Username    string `json:"username"`
	Nationality string `json:"nationality"`
	Ethnicity   string `json:"ethnicity"`
	Telephone   string `json:"telephone"`
}
