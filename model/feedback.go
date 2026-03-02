package model

import "time"

type Feedback struct {
	Oid              int64     `json:"oid"`
	ByUsername       string    `json:"byUsername"`
	RatingDate       time.Time `json:"ratingDate"`
	Rating           string    `json:"rating"`
	Disputed         bool      `json:"disputed"`
	Feedback         string    `json:"feedback"`
	FeedbackResponse string    `json:"feedbackResponse"`
}
