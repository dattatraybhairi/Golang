package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes" //add this

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(app) //add this

	app.Listen(":3000")
}
