package main

import (
	"fmt"
	"net/http"
	"phoenix-client-service/handler"
	"phoenix-client-service/ws"
)

func main() {

	// Run the websocket hub (for injecting message into)
	go ws.Hub.Run()

	http.HandleFunc("/ping", handler.HandlePingRequest)
	http.HandleFunc("/search", handler.HandleSearchRequest)
	http.HandleFunc("/feedback", handler.HandleFeedbackRequest)
	http.HandleFunc("/servicereports", handler.HandleServiceReportRequest)
	http.HandleFunc("/watchlist", handler.AddUserIdToWatchList)
	http.HandleFunc("/watchlist/all", handler.GetAllWatchListEntries)
	http.HandleFunc("/watchlist/today", handler.GetTodaysWatchListEntries)
	http.HandleFunc("/watchlist/import", handler.HandleImportWatchListFeed)
	http.HandleFunc("/orders", handler.HandleSubmitOrder)
	http.HandleFunc("/orders/all", handler.HandleGetAllOrders)
	http.HandleFunc("/orders/import", handler.HandleImportOrderFeed)
	http.HandleFunc("/orders/filter", handler.HandleGetAllOrdersByYear)
	http.HandleFunc("/orders/generateRequest", handler.HandleGenerateOrderRequest)
	http.HandleFunc("/orders/cancel", handler.HandleCancelOrder)
	http.HandleFunc("/link", handler.HandleAddLink)
	http.HandleFunc("/link/remove", handler.HandleRemoveLink)
	http.HandleFunc("/monitor/ping", handler.PingHandler)
	http.HandleFunc("/ws", handler.WsHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error in starting server", err)
	}
}
