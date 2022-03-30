package eventhandler

import (
	"encoding/json"

	"wptest/services.customer/src/internal"
	"wptest/services.customer/src/entity"
)

func CreateProduct(service *customer.Service, message []byte) {
	var product entity.Product
	json.Unmarshal(message, &product)
	service.CreateProduct(product)
}
