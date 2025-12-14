package handler

import (
	"encoding/json"
	"net/http"
	"phoenix-client-service/dao"
	"phoenix-client-service/model"
)

func HandleAddLink(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "POST" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a POST")
		return
	}

	var linkRequest model.LinkRequest
	err := json.NewDecoder(req.Body).Decode(&linkRequest)

	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}

	result, err := dao.AddLink(linkRequest.UserId1, linkRequest.UserId2, linkRequest.InputType, linkRequest.Notes)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	if result {
		res.WriteHeader(201)
		res.Header().Add("Content-Type", "application/json")
	} else {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", "Unable to save link.")
		return
	}
}

func HandleRemoveLink(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "DELETE" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a DELETE")
		return
	}

	var linkRequest model.LinkRequest
	err := json.NewDecoder(req.Body).Decode(&linkRequest)

	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}

	result, err := dao.RemoveLink(linkRequest.UserId1, linkRequest.UserId2)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	if result {
		res.WriteHeader(200)
		res.Header().Add("Content-Type", "application/json")
	} else {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", "Unable to save link.")
		return
	}
}
