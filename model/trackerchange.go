package model

type TrackerChange struct {
	Oid            int64  `json:"oid"`
	Username       string `json:"username"`
	UserId         string `json:"userId"`
	FieldValue     string `json:"fieldValue"`
	OldValue       string `json:"oldValue"`
	NewValue       string `json:"newValue"`
	RecordDateTime string `json:"recordDateTime"`
}
