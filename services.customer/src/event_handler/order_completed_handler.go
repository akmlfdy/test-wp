package eventhandler

import (
	"encoding/json"

	"wptest/services.customer/src/event"
	"wptest/services.customer/src/internal"
)

func ClearBasket(service *customer.Service, message []byte) {
	var order event.OrderCreated
	json.Unmarshal(message, &order)
	basket, _ := service.GetBasket(order.CustomerID)
	service.ClearBasketItems(basket.ID)
}
