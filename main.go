package main

import (
	"net/http"
	"phoenix-client-service/handler"
)

func main() {

	http.HandleFunc("/ping", handler.HandlePingRequest)
	http.HandleFunc("/search", handler.HandleSearchRequest)
	http.HandleFunc("/feedback", handler.HandleFeedbackRequest)
	http.ListenAndServe(":8081", nil)
}
