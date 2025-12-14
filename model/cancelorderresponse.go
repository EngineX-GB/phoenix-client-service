package model

type CancelOrderResponse struct {
	OrderId     uint64 `json:"orderId"`
	OrderStatus string `json:"orderStatus"`
}
