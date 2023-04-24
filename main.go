package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"proxy-handler/config"
	"proxy-handler/notifications"
)

func main() {
	// load env variables
	config.LoadEnv()

	// load config

	app := fiber.New()

	router := app.Group("/notifications")

	notifications.RegisterRoutes(router)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
