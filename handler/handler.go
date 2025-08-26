package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"phoenix-client-service/dao"
	"phoenix-client-service/model"
	"time"
)

/**/
func HandlePingRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Ping")
}

func HandleFeedbackRequest(res http.ResponseWriter, req *http.Request) {
	// do validation on the method to check if the method is a get of not
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
	if req.Method != "GET" {
		var entry model.ErrorResponse
		entry.Timestamp = time.Now()
		entry.Message = "405 Method Not Supported"
		entry.Description = "Method must be be a Post"
		entry.ServiceName = "phoenix-client-service"
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(405)
		data, err := json.Marshal(entry)
		if err != nil {
			//handle error
			log.Fatal("Unable to form a response, due to JSON marshalling error")
		}
		res.Write(data)
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
			log.Panic("Error in retrieving entries from the database")
		}

		data, err := json.Marshal(entries)
		if err != nil {
			log.Panic("Unable to marshal entries into JSON format")
		}

		res.Write(data)
		res.WriteHeader(200)
		res.Header().Add("Content-Type", "application/json")
	}

}

func HandleTrackerChangeRequest(res http.ResponseWriter, req *http.Request) {}
