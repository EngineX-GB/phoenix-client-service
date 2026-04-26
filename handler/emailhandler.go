package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"phoenix-client-service/model"
	"phoenix-client-service/service"
	"strconv"
)

func AddEmailConfig(res http.ResponseWriter, req *http.Request) {
	var emailConfig model.SMTPConfig
	var errorResponse model.ErrorResponse

	if req.Method != "POST" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a POST")
		return
	}

	err := json.NewDecoder(req.Body).Decode(&emailConfig)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}
	err = service.AddEmailConfig(emailConfig)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(201)
}

func AddEmailAccount(res http.ResponseWriter, req *http.Request) {

	var emailAccount model.EmailAccount
	var errorResponse model.ErrorResponse

	if req.Method != "POST" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a POST")
		return
	}

	err := json.NewDecoder(req.Body).Decode(&emailAccount)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}
	err = service.AddEmailAccount(emailAccount)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(201)
}

func DeleteEmailConfig(res http.ResponseWriter, req *http.Request) {
	var oid_string = req.URL.Query().Get("id")
	var errorResponse model.ErrorResponse

	if req.Method != "DELETE" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a DELETE")
		return
	}

	oid, err := strconv.ParseInt(oid_string, 10, 64)

	if err != nil {
		log.Printf("An error has occurred when processing the request: %s", err.Error())
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}

	err = service.DeleteEmailConfig(oid)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(200)
}

func DeleteEmailAccount(res http.ResponseWriter, req *http.Request) {
	var oid_string = req.URL.Query().Get("id")
	var errorResponse model.ErrorResponse

	if req.Method != "DELETE" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a DELETE")
		return
	}

	oid, err := strconv.ParseInt(oid_string, 10, 64)

	if err != nil {
		log.Printf("An error has occurred when processing the request: %s", err.Error())
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}

	err = service.DeleteEmailAccount(oid)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(200)
}

func GetEmailConfigs(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "GET" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a GET")
		return
	}
	listOfConfigs, err := service.GetAllEmailConfigs()
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}
	data, err := json.Marshal(listOfConfigs)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(200)
	res.Write(data)
}

func GetEmailAccounts(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "GET" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a GET")
		return
	}
	listOfAccounts, err := service.GetAllEmailAccounts()
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}
	data, err := json.Marshal(listOfAccounts)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(200)
	res.Write(data)
}

func HandleSendEmail(res http.ResponseWriter, req *http.Request) {
	var emailRequest model.EmailRequest
	var errorResponse model.ErrorResponse
	err := json.NewDecoder(req.Body).Decode(&emailRequest)
	if err != nil {
		// handle the error with a 500 internal error
		fmt.Println(err.Error())
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}

	// send the email here
	err = service.SendEmail(emailRequest.EmailAccountId, emailRequest.ToEmailAddress, emailRequest.Subject, emailRequest.Message)

	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(200)
}
