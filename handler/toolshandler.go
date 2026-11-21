package handler

import (
	"encoding/json"
	"net/http"
	"phoenix-client-service/model"
	"phoenix-client-service/service"
)

func TelephoneSearchHandler(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "GET" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a GET")
		return
	}

	queryParam := req.URL.Query().Get("searchQuery")

	if queryParam == "" {
		errorResponse.PublishErrorResponse(res, 400, "Bad Request", "query parameter 'searchQuery' cannot be blank or null")
		return
	}

	records, err := service.SearchTelephoneNumber(queryParam)

	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}

	recordsJson, err := json.Marshal(records)

	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	res.WriteHeader(200)
	res.Header().Add("Content-Type", "application/json")
	res.Write(recordsJson)

}

func UsernameSearchHandler(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "GET" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a GET")
		return
	}

	queryParam := req.URL.Query().Get("searchQuery")

	if queryParam == "" {
		errorResponse.PublishErrorResponse(res, 400, "Bad Request", "query parameter 'searchQuery' cannot be blank or null")
		return
	}

	records, err := service.SearchUsername(queryParam)

	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}

	recordsJson, err := json.Marshal(records)

	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	res.WriteHeader(200)
	res.Header().Add("Content-Type", "application/json")
	res.Write(recordsJson)

}
