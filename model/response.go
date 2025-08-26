package model

import "time"

type ErrorResponse struct {
	ServiceName string    `json:"serviceName"`
	Timestamp   time.Time `json:"timestamp"`
	Message     string    `json:"message"`
	Description string    `json:"description"`
}
