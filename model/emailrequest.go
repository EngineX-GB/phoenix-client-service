package model

type EmailRequest struct {
	EmailAccountId int64  `json:"emailAccountId"`
	ToEmailAddress string `json:"toEmailAddress"`
	Subject        string `json:"subject"`
	Message        string `json:"message"`
}
