package eventhandler

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	gouuid "github.com/satori/go.uuid"
	"wptest/services.order/src/entity"
	"wptest/services.order/src/internal"
	"wptest/shared/config"
	"wptest/shared/kafka"
)

func FailOrder(service *order.Service, messageKey []byte) {

	orderID, _ := gouuid.FromString(string(messageKey))
	order, err := service.UpdateOrderStatus(uuid.UUID(orderID), int(entity.OrderFailed))
	if err != nil {
		log.Printf("FailOrder.UpdateOrderStatus failed: %v", err.Error())
		return
	}
	payload, _ := json.Marshal(order)
	kafka.Publish(order.ID, payload, "OrderFailed", config.AppConfig.KafkaOrderTopic)
}
