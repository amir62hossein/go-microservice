package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {

	app := fiber.New()

	// auth proxy (login, register)
	app.Post("/login", proxy.Forward("http://localhost:8001/login"))
	app.Post("/register", proxy.Forward("http://localhost:8001/register"))

	// product proxy (buy, create)
	app.Post("/product/buy/:id", func(c *fiber.Ctx) error {
		return proxy.Forward("http://localhost:8002/product/buy/" + c.Params("id"))(c)
	})
	app.Post("/product/create", proxy.Forward("http://localhost:8002/product/create"))
	app.Get("/products", proxy.Forward("http://localhost:8002/products"))

	log.Fatal(app.Listen(":8000"))
}
