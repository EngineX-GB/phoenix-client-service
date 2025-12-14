package handler

import (
	"encoding/json"
	"net/http"
	"phoenix-client-service/dao"
	"phoenix-client-service/model"
	"strconv"
)

func HandleGetAllOrders(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "GET" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a GET")
		return
	}
	orders, err := dao.GetAllOrders()
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	data, err := json.Marshal(orders)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(200)
	res.Write(data)
}

func HandleSubmitOrder(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	var requestPayload model.Order
	if req.Method != "POST" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a POST")
		return
	}
	// otherwise, process the order

	err := json.NewDecoder(req.Body).Decode(&requestPayload)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	//put the record in the database.
	err = dao.SaveOrder(requestPayload)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", "An error occurred when saving the order : "+err.Error())
		return
	}
	res.WriteHeader(201)
	res.Header().Add("Content-Type", "application/json")
}

func HandleCancelOrder(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "PUT" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a PUT")
		return
	}
	orderId := req.URL.Query().Get("id")
	orderIdUint, err := strconv.ParseUint(orderId, 10, 64)

	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}

	result, err := dao.UpdateOrderStatus(orderIdUint, "CANCELLED")
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	var cancelOrderResponse model.CancelOrderResponse

	if result {
		cancelOrderResponse.OrderId = orderIdUint
		cancelOrderResponse.OrderStatus = "CANCELLED"
		res.WriteHeader(200)
		res.Header().Add("Content-Type", "application/json")
		data, err := json.Marshal(cancelOrderResponse)
		if err != nil {
			panic(err)
		}
		res.Write(data)
	} else {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", "The order status was not updated.")
		return
	}

}

func HandleGenerateOrderRequest(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "GET" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a GET")
		return
	}
	userId := req.URL.Query().Get("userId")
	orderRequest, err := dao.RetrieveOrderRequestDetails(userId)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	data, err := json.Marshal(orderRequest)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	res.WriteHeader(200)
	res.Header().Add("Content-Type", "application/json")
	res.Write(data)
}

func HandleGetAllOrdersByYear(res http.ResponseWriter, req *http.Request) {
	var errorResponse model.ErrorResponse
	if req.Method != "GET" {
		errorResponse.PublishErrorResponse(res, 405, "Method Not Allowed", "Method must be a GET")
		return
	}

	queryParameterYear := req.URL.Query().Get("year")
	year, err := strconv.Atoi(queryParameterYear)
	if err != nil {
		panic(err)
	}
	orders, err := dao.GetAllOrdersWithFilter(year)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	data, err := json.Marshal(orders)
	if err != nil {
		errorResponse.PublishErrorResponse(res, 500, "Internal Server Error", err.Error())
		return
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(200)
	res.Write(data)
}
