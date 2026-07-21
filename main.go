package main

import (
	"fmt"
	"net/http"
	"phoenix-client-service/handler"
	"phoenix-client-service/ws"
	"runtime"
)

func main() {

	fmt.Println("Running Phoenix Client Service for " + runtime.GOOS)

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
	http.HandleFunc("/notifications/publish", handler.HandleNotificationEmitter)
	http.HandleFunc("/comms/email/send", handler.HandleSendEmail)
	http.HandleFunc("/comms/email/configs", handler.GetEmailConfigs)
	http.HandleFunc("/comms/email/accounts", handler.GetEmailAccounts)
	http.HandleFunc("/comms/email/accounts/add", handler.AddEmailAccount)
	http.HandleFunc("/comms/email/config/add", handler.AddEmailConfig)
	http.HandleFunc("/comms/email/account/remove", handler.DeleteEmailAccount)
	http.HandleFunc("/comms/email/config/remove", handler.DeleteEmailConfig)
	http.HandleFunc("/market/load/region", handler.LoadMarketByRegion)
	http.HandleFunc("/ws", handler.WsHandler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error in starting server", err)
	}

}
