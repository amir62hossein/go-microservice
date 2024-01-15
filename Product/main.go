package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"product-service/handler"
	"product-service/middleware"
)

func main() {
	_ = godotenv.Load(".env")
	app := fiber.New()

	app.Post("/product/buy/:id", middleware.Authentication(), handler.BuyProduct)
	app.Post("/product/create", handler.CreateProduct)
	app.Get("/products", handler.GetAllProducts)

	log.Fatal(app.Listen(":8002"))
}
