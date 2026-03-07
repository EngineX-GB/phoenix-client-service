package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"phoenix-client-service/dao"
	"phoenix-client-service/model"
)

func HandleServiceReportRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		var entry model.ErrorResponse
		entry.PublishErrorResponse(res, 405, "Method Not Supported", "Method must be a GET")
	}
	userId := req.URL.Query().Get("userId")

	entries, err := dao.ExecuteServiceReportHeadlineQuery(userId)
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
