package main

import (
	"net/http"
	"phoenix-client-service/handler"
)

func main() {

	http.HandleFunc("/ping", handler.HandlePingRequest)
	http.HandleFunc("/search", handler.HandleSearchRequest)
	http.HandleFunc("/feedback", handler.HandleFeedbackRequest)
	http.HandleFunc("/watchlist", handler.AddUserIdToWatchList)
	http.HandleFunc("/watchlist/all", handler.GetAllWatchListEntries)
	http.HandleFunc("/watchlist/today", handler.GetTodaysWatchListEntries)
	http.HandleFunc("/orders", handler.HandleSubmitOrder)
	http.HandleFunc("/orders/all", handler.HandleGetAllOrders)
	http.HandleFunc("/orders/filter", handler.HandleGetAllOrdersByYear)
	http.ListenAndServe(":8081", nil)
}
