package model

import "time"

type ErrorResponse struct {
	ServiceName string    `json:"serviceName"`
	Timestamp   time.Time `json:"timestamp"`
	Message     string    `json:"message"`
	Description string    `json:"description"`
}

func (e ErrorResponse) PublishErrorResponse(message string, description string) ErrorResponse {
	var response ErrorResponse
	response.ServiceName = "phoenix-client-service"
	response.Description = description
	response.Message = message
	response.Timestamp = time.Now()
	return response
}
