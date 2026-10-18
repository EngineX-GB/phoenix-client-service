package handler

import (
	"context"
	"fmt"
	"net/http"
	"phoenix-client-service/model"
	"phoenix-client-service/service"
	"sync"
)

var (
	serviceNames = map[string]bool{
		"market":    false,
		"ingestion": false,
	}

	// Stores cancel functions for each running pinger
	pingers = map[string]context.CancelFunc{}

	// Mutex to protect both maps
	mu sync.Mutex
)

func PingHandler(res http.ResponseWriter, req *http.Request) {
	var serviceName string = req.URL.Query().Get("serviceName")
	var errorResponse model.ErrorResponse

	if serviceName == "" {
		errorResponse.PublishErrorResponse(res, 400, "Bad Request", "Missing service name in the query parameters")
		return
	}

	mu.Lock()
	value, exists := serviceNames[serviceName]
	if !exists {
		mu.Unlock()
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", "The service queried is not valid")
		return
	}
	if value {
		// Service is ON → turn OFF
		fmt.Printf("Turning OFF service: %s\n", serviceName)

		// Stop goroutine
		if stopFunc, ok := pingers[serviceName]; ok {
			stopFunc()
			delete(pingers, serviceName)
		}

		serviceNames[serviceName] = false
	} else {
		// Service is OFF → turn ON
		fmt.Printf("Turning ON service: %s\n", serviceName)

		// Start goroutine
		var serverURL string
		if serviceName == "market" {
			serverURL = service.GetMarketConnectorEndpoint()
		}
		if serviceName == "ingestion" {
			serverURL = service.GetIngestionServiceEndpoint()
		}

		stopFunc := service.StartPinger(serviceName, serverURL)
		pingers[serviceName] = stopFunc

		serviceNames[serviceName] = true
	}

	mu.Unlock()

	fmt.Fprintf(res, "Service %s is now %v\n", serviceName, serviceNames[serviceName])

}
