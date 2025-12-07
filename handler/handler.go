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

func GetAllWatchListEntries(res http.ResponseWriter, req *http.Request) {
	var entry model.ErrorResponse
	if req.Method != "GET" {
		entry.PublishErrorResponse(res, 405, "Method Not Supported", "Method must be a GET")
		return
	}
	entries, err := dao.ExecuteGetWatchlist(false)
	if err != nil {
		log.Printf("Unable to get watchlist data")
		log.Panic(err)
	}
	data, err := json.Marshal(entries)
	if err != nil {
		log.Printf("Unable to marshal entries into Json")
		log.Panic(err)
	}
	res.WriteHeader(200)
	res.Header().Add("Content-Type", "application/json")
	res.Write(data)
}

func GetTodaysWatchListEntries(res http.ResponseWriter, req *http.Request) {
	var entry model.ErrorResponse
	if req.Method != "GET" {
		entry.PublishErrorResponse(res, 405, "Method Not Supported", "Method must be a GET")
		return
	}
	entries, err := dao.ExecuteGetWatchlist(true)
	if err != nil {
		log.Printf("Unable to get watchlist data")
		log.Panic(err)
	}
	data, err := json.Marshal(entries)
	if err != nil {
		log.Printf("Unable to marshal entries into Json")
		log.Panic(err)
	}
	res.WriteHeader(200)
	res.Header().Add("Content-Type", "application/json")
	res.Write(data)
}

func AddUserIdToWatchList(res http.ResponseWriter, req *http.Request) {
	var entry model.ErrorResponse
	var watchListRequest model.WatchListRequest
	if req.Method != "POST" {
		entry.PublishErrorResponse(res, 405, "Method Not Supported", "Method must be a POST")
		return
	}
	err := json.NewDecoder(req.Body).Decode(&watchListRequest)
	if err != nil {
		entry.PublishErrorResponse(res, 500, "Error", err.Error())
		return
	}
	newRecordId, err := dao.ExecuteAddWatchListEntry(watchListRequest.UserId)
	if err != nil {
		entry.PublishErrorResponse(res, 500, "Error", err.Error())
		return
	}
	if newRecordId == -1 {
		entry.PublishErrorResponse(res, 500, "Error", "An error has occured when trying to add a watchlist entry")
		return
	}
	// otherwise it's a valid response
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(int(201))
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
