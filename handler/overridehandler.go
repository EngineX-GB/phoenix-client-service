package handler

import (
	"encoding/json"
	"net/http"
	"phoenix-client-service/model"
	"phoenix-client-service/service"
)

// functions:
// get all overrides - done
// add new override for user - done
// update existing override for user
// delete an existing override

func test() string {
	return ""
}

func HandleGetAllOverridesRequest(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "GET" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Request must be a GET")
		return
	}
	overrideData, err := service.GetAllDataOverrides()

	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}

	data, err := json.Marshal(overrideData)

	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	res.WriteHeader(200)
	res.Header().Add("Content-Type", "application/json")
	res.Write(data)
}

func HandleDeleteOverrideRequest(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "DELETE" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Request must be a DELETE")
		return
	}

	userId := req.URL.Query().Get("userId")       // the oid in the database
	fieldToDelete := req.URL.Query().Get("field") // e.g. telephone, email, etc

	err := service.DeleteDataOverride(userId, fieldToDelete)

	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}

}

func HandleAddOverrideRequest(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "POST" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Request must be a POST")
		return
	}
	var requestPayload model.OverrideReqeust

	err := json.NewDecoder(req.Body).Decode(&requestPayload)

	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	err = service.AddNewOverride(requestPayload.UserId, requestPayload.Data)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}
	res.WriteHeader(201)
	res.Header().Add("Content-Type", "application/json")
}
