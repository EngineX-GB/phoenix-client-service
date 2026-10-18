package handler

import (
	"encoding/json"
	"net/http"
	"phoenix-client-service/model"
	"phoenix-client-service/service"
	"time"
)

func LoadMarketByRegion(res http.ResponseWriter, req *http.Request) {

	var errorResponse model.ErrorResponse
	if req.Method != "POST" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a POST")
		return
	}

	regionId := req.URL.Query().Get("id")

	if regionId == "" {
		errorResponse.PublishErrorResponse(res, 400, "Bad Request", "No 'id' query parameter has been provided")
		return
	}

	// perform the rest call to the market connector

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest(http.MethodPost, service.GetMarketConnectorEndpoint()+"/load/region/"+regionId, nil)
	if err != nil {
		print(err.Error())
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		print(err.Error())
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}
	defer resp.Body.Close()

	var loadResp model.MarketLoadResponse
	err = json.NewDecoder(resp.Body).Decode(&loadResp)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Service Error", err.Error())
		return
	}

	data, err := json.Marshal(loadResp)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(201)
	res.Write(data)
}
