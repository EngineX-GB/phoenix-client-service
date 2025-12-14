package dao

import (
	"phoenix-client-service/datasource"
	"phoenix-client-service/model"
	"phoenix-client-service/util"
)

func UpdateOrderStatus(orderId uint64, orderStatus string) (bool, error) {
	db := datasource.Connect()
	tx, err := db.Begin()
	defer tx.Rollback()
	defer db.Close()

	if err != nil {
		print("An error occured while trying to update an order : " + err.Error())
		return false, err
	}

	_, err = tx.Exec(util.UpdateOrderStatus(), orderStatus, orderId)

	if err != nil {
		print("An error occured while trying to update an order : " + err.Error())
		return false, err
	}

	if err := tx.Commit(); err != nil {
		print("An error occured while trying to update an order : " + err.Error())
		return false, err
	}
	return true, nil
}

func SaveOrder(order model.Order) error {
	db := datasource.Connect()
	tx, err := db.Begin()
	defer tx.Rollback() // defer the rollback so it's always called if something goes wrong
	defer db.Close()

	if err != nil {
		print("An error occured while trying to save an order : " + err.Error())
		return err
	}

	_, err = tx.Exec(util.SaveOrderEntry(), order.UserId, order.UserName, order.Location, order.Region, order.DateOfEvent,
		order.TimeOfEvent, order.CreationDate, order.ModifiedDate, order.Duration, order.Rate, order.Deductions,
		order.Surplus, order.Price, order.Status, order.Notes)

	if err != nil {
		print("aAn error has occurred when trying to save an order : " + err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		print("An error has occurred when trying to save an order : " + err.Error())
		return err
	}

	return nil
}

func GetAllOrders() ([]model.Order, error) {
	return GetAllOrdersWithFilter(-1)
}

func RetrieveOrderRequestDetails(userId string) (model.OrderRequest, error) {
	var orderRequest model.OrderRequest
	db := datasource.Connect()
	rows, err := db.Query(util.GenerateOrderRequest(userId))
	if err != nil {
		return orderRequest, err
	}
	defer db.Close()
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&orderRequest.UserId, &orderRequest.UserName, &orderRequest.Location,
			&orderRequest.Region, &orderRequest.Telephone, &orderRequest.R15, &orderRequest.R30, &orderRequest.R45,
			&orderRequest.R100, &orderRequest.R150, &orderRequest.R200, &orderRequest.R250,
			&orderRequest.R300, &orderRequest.R350, &orderRequest.R400, &orderRequest.RON); err != nil {
			print(err.Error())
		}
	}
	return orderRequest, nil
}

func GetAllOrdersWithFilter(year int) ([]model.Order, error) {
	var orders = make([]model.Order, 0)
	db := datasource.Connect()

	var queryString string
	if year == -1 {
		queryString = util.GetAllOrders()
	} else {
		queryString = util.GetOrdersByYear(year)
	}

	rows, err := db.Query(queryString)

	if err != nil {
		print("An error has occurred when trying to retrieve orders : " + err.Error())
		return nil, err
	}
	defer rows.Close()
	defer db.Close()
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.UserId, &order.UserName, &order.Location, &order.Region, &order.DateOfEvent, &order.TimeOfEvent,
			&order.Duration, &order.Rate, &order.Deductions, &order.Surplus,
			&order.Price, &order.Status, &order.Notes); err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return orders, err
	}
	return orders, nil
}
