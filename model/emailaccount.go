package model

type EmailAccount struct {
	Oid          int64  `json:"oid"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	FormalName   string `json:"formalName"`
	SmtpConfigId int64  `json:"emailConfigId"`
}
