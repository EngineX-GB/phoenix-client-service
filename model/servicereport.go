package model

import "time"

type ServiceReportHeadline struct {
	Oid          int64     `json:"oid"`
	UserId       string    `json:"userId"`
	MeetDate     time.Time `json:"meetDate"`
	ReportRating string    `json:"reportRating"`
	HeadLine     string    `json:"headline"`
}
