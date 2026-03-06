package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"phoenix-client-service/model"
	"phoenix-client-service/service"
)

/*
*
POST method to import the watchlist feed into the database.
*/
func HandleImportWatchListFeed(res http.ResponseWriter, req *http.Request) {

	var errorResponse model.ErrorResponse
	if req.Method != "POST" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a POST")
		return
	}

	file, metadata, err := req.FormFile("file")
	if err != nil {
		fmt.Println("Error in reading uploaded file : ", err)
		return
	}
	defer file.Close()
	log.Println("Loading file : " + metadata.Filename)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error in reading uploaded file : ", err)
		return
	}
	contents := string(fileBytes[:])

	//TODO: Read the file here and process it 06-03-2026
	service.ReadWatchListFeed(contents)
}
