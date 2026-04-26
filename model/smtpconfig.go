package model

type SMTPConfig struct {
	Oid  int64  `json:"oid"`
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}
