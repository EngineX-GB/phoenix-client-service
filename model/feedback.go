package model

import "time"

type Feedback struct {
	Oid             int64     `json:"oid"`
	ServiceProvider string    `json:"serviceProvider"`
	Rating          string    `json:"rating"`
	Comment         string    `json:"comment"`
	FeedbackDate    time.Time `json:"feedbackDate"`
}
