package main

import (
	"auth-service/handler"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}

	app := fiber.New()

	app.Post("/login", handler.Login)
	app.Post("/register", handler.Register)

	log.Fatal(app.Listen(":8001"))
}
