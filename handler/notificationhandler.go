package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"phoenix-client-service/model"
	"phoenix-client-service/ws"
)

func HandleNotificationEmitter(res http.ResponseWriter, req *http.Request) {
	var entry model.ErrorResponse
	var notificationPayload model.Notification
	if req.Method != "POST" {
		entry.PublishErrorResponse(res, 405, "Method Not Supported", "Method must be a POST")
		return
	}
	err := json.NewDecoder(req.Body).Decode(&notificationPayload)
	if err != nil {
		entry.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		fmt.Println("[DEBUG] Error in publishing notification : " + err.Error())
		return
	}

	// broadcast the notification

	fmt.Println("[DEBUG] Publishing notification payload to clients")
	jsonBytes, _ := json.Marshal(notificationPayload)
	ws.Hub.Broadcast <- jsonBytes

	res.WriteHeader(201)
	res.Header().Add("Content-Type", "application/json")
}
