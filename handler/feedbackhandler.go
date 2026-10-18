package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"phoenix-client-service/model"
	"phoenix-client-service/service"
)

func HandleLoadFeedbackRequest(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "POST" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a POST")
		return
	}

	// the request should contain a collection of strings (i.e. userids)

	var userIdList []string

	err := json.NewDecoder(req.Body).Decode(&userIdList)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		fmt.Println("[DEBUG] Error in forwarding request to market connector : " + err.Error())
		return
	}

	body, err := json.Marshal(userIdList)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		fmt.Println("[DEBUG] Error in forwarding request to market connector : " + err.Error())
		return
	}

	newReq, err := http.NewRequest(http.MethodPost, service.GetMarketConnectorEndpoint()+"/load/feedbackv2", bytes.NewBuffer(body))
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		fmt.Println("[DEBUG] Error in forwarding request to market connector : " + err.Error())
		return
	}

	newReq.Header.Set("Content-Type", "application/json")

	// send the request

	resp, err := http.DefaultClient.Do(newReq)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		fmt.Println("[DEBUG] Error in forwarding request to market connector : " + err.Error())
		return
	}
	defer resp.Body.Close()

	//TODO: Handle the failure message for when the response from the market connector is not OK

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(202)

}
