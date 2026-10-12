package model

type OverrideReqeust struct {
	UserId string       `json:"userId"`
	Data   OverrideData `json:"data"`
}

type OverrideData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type OverrideEntity struct {
	Oid      int64          `json:"oid"`
	UserId   string         `json:"userId"`
	UserName string         `json:"username"`
	Data     []OverrideData `json:"data"`
}

type OverrideRecord struct {
	Oid      int64  `json:"oid"`
	UserId   string `json:"userId"`
	UserName string `json:"username"`
	Data     string `json:"data"`
}

type SingleOverrideRecord struct {
	Oid           int64  `json:"oid"`
	UserId        string `json:"userId"`
	UserName      string `json:"username"`
	PropertyName  string `json:"propertyName"`
	PropertyValue string `json:"propertyValue"`
}
