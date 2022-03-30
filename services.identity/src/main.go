package main

import (
	"wptest/services.identity/src/entity"
	identity "wptest/services.identity/src/internal"
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

	db.AutoMigrate(&entity.User{})

	repo := identity.NewRepository(db)
	service := identity.NewService(repo)
	handler := identity.NewHandler(service)

	err = server.NewServer(handler.Init(), config.AppPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}
}
