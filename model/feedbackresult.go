package model

type FeedbackResult struct {
	Headers      QueryMarker `json:"headers"`
	FeedbackData []Feedback  `json:"feedbackData"`
}
