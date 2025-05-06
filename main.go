package main

import (
	"wallet-app/config"
	"wallet-app/models"
	"wallet-app/routes"
)

func seedUsers() {
	var count int64
	config.DB.Model(&models.User{}).Count(&count)

	if count == 0 {
		users := []models.User{
			{Name: "deddy1", Balance: 1000000},
			{Name: "deddy2", Balance: 5000000},
		}
		for _, user := range users {
			config.DB.Create(&user)
		}
	}
}

func main() {
	config.InitDB()
	config.DB.AutoMigrate(&models.User{})
	seedUsers()
	router := routes.SetupRouter()
	router.Run(":8080")
}
