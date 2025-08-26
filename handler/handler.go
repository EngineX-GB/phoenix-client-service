package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"phoenix-client-service/dao"
	"phoenix-client-service/model"
)

/**/
func HandlePingRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Ping")
}

func HandleFeedbackRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		var entry model.ErrorResponse
		entry.PublishErrorResponse(res, 405, "Method Not Supported", "Method must be a GET")
	}
	userId := req.URL.Query().Get("userId")
	entries, err := dao.ExecuteFeedbackQuery(userId)
	if err != nil {
		log.Panic(err)
	}
	data, err := json.Marshal(entries)
	if err != nil {
		log.Printf("Unable to marshal entries into json")
		log.Panic(err)
	}
	res.WriteHeader(200)
	res.Header().Add("Content-Type", "application/json")
	res.Write(data)
}

func HandleSearchRequest(res http.ResponseWriter, req *http.Request) {
	var entry model.ErrorResponse
	if req.Method != "GET" {
		entry.PublishErrorResponse(res, 405, "Method Not Supported", "Method must be a GET")
	} else {
		// run the query
		username := req.URL.Query().Get("username")
		nationality := req.URL.Query().Get("nationality")
		userId := req.URL.Query().Get("userId")
		region := req.URL.Query().Get("region")

		entries, err := dao.ExecuteSearchQuery(model.SearchRequest{Username: username,
			Nationality: nationality,
			UserId:      userId,
			Region:      region})

		if err != nil {
			log.Println("[ERROR] Error in retrieving entries from the database. Error: " + err.Error())
			entry.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
			return
		}

		data, err := json.Marshal(entries)
		if err != nil {
			log.Panic("Unable to marshal entries into JSON format. Error: " + err.Error())
			entry.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		}

		res.Write(data)
		res.Header().Add("Content-Type", "application/json")
	}

}

func HandleTrackerChangeRequest(res http.ResponseWriter, req *http.Request) {}
