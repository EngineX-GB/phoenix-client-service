package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"phoenix-client-service/config"
	"phoenix-client-service/model"
)

func ConfigHandler(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "GET" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Request must be a GET")
		return
	}

	configMap := config.GetListOfConfigs()

	data, err := json.Marshal(configMap)
	if err != nil {
		log.Fatal("Unable to form a response, due to JSON marshalling error. Error " + err.Error())
	}

	res.WriteHeader(200)
	res.Header().Add("Content-Type", "application/json")
	res.Write(data)
}
