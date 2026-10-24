package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"phoenix-client-service/handler"
	"phoenix-client-service/model"
	"phoenix-client-service/service"
	"phoenix-client-service/ws"
	"runtime"
	"text/tabwriter"
)

// go will embed version.json. The directive below will reference the version.json

//go:embed version.json
var versionJSON []byte

func envDump() {
	fmt.Println()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, "PROPERTY\tVALUE")
	fmt.Fprintln(w, "--------\t-----")
	fmt.Fprintf(w, "Environment\t%s\n", service.GetEnvironment())
	fmt.Fprintf(w, "Datasource Host\t%s\n", service.GetDataSourceHostname())
	fmt.Fprintf(w, "Datasource Name\t%s\n", service.GetDataSourceName())
	fmt.Fprintf(w, "Ingestion Service Endpoint\t%s\n", service.GetIngestionServiceEndpoint())
	fmt.Fprintf(w, "Market Connector Endpoint\t%s\n", service.GetMarketConnectorEndpoint())

	w.Flush()
	fmt.Println()

}

func getVersion() string {

	var v model.Version

	if err := json.Unmarshal(versionJSON, &v); err != nil {
		panic(err)
	}
	return v.VersionNumber
}

func main() {

	fmt.Println("Running Phoenix Client Service " + getVersion() + " for " + runtime.GOOS)

	envDump()

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
	http.HandleFunc("/feedback/load", handler.HandleLoadFeedbackRequest)
	http.HandleFunc("/override/add", handler.HandleAddOverrideRequest)
	http.HandleFunc("/override/all", handler.HandleGetAllOverridesRequest)
	http.HandleFunc("/override/remove", handler.HandleDeleteOverrideRequest)
	http.HandleFunc("/ws", handler.WsHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%s", service.GetServicePort()), nil)
	if err != nil {
		fmt.Println("Error in starting server", err)
	}

}
