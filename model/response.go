package model

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ErrorResponse struct {
	ServiceName string    `json:"serviceName"`
	Timestamp   time.Time `json:"timestamp"`
	Message     string    `json:"message"`
	Description string    `json:"description"`
}

func (e ErrorResponse) PublishErrorResponse(res http.ResponseWriter, statusCode int, message string, description string) {
	var errorResponse ErrorResponse
	errorResponse.ServiceName = "phoenix-client-service"
	errorResponse.Description = description
	errorResponse.Message = message
	errorResponse.Timestamp = time.Now()
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(int(statusCode))
	data, err := json.Marshal(errorResponse)
	if err != nil {
		log.Fatal("Unable to form a response, due to JSON marshalling error. Error " + err.Error())
	}
	res.Write(data)
}
