package main

import (
	"wptest/services.order/src/entity"
	order "wptest/services.order/src/internal"
	"wptest/services.order/src/kafka"
	"wptest/shared/config"
	"wptest/shared/server"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	config := config.LoadConfig(".")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.GetDBURL(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the DB.")
	}

	db.AutoMigrate(&entity.Order{})
	db.AutoMigrate(&entity.OrderItem{})

	repo := order.NewRepository(db)
	service := order.NewService(repo)
	handler := order.NewHandler(service)

	go kafka.RegisterConsumer(config.KafkaProductTopic, service)

	err = server.NewServer(handler.Init(), config.AppPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}
}
