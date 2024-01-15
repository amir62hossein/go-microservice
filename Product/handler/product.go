package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"product-service/database"
	"product-service/helper"
	"product-service/model"
	"product-service/rabbitMQ"
	"time"
)

func BuyProduct(c *fiber.Ctx) error {
	broker := rabbitMQ.NewRabbitMQ()
	db := database.NewPostgres()
	productId, err := c.ParamsInt("id")
	userID := c.Locals("userID").(string)
	userEmail := c.Locals("userEmail").(string)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var product model.Product
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": err.Error()})
	}
	if err := db.Db.Where("id = ?", productId).Find(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}

	productMessage := helper.MessageConstructor(userID, userEmail, product.Name)

	orderQueue, _ := broker.Channel.QueueDeclare("ORDER",
		true,
		false,
		false,
		false,
		nil)
	err = broker.Channel.PublishWithContext(ctx,
		"",
		orderQueue.Name,
		false,false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        productMessage,
		})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}
	log.Println("[ # ] Product Add To Order Queue")


	return nil
}
func CreateProduct(c *fiber.Ctx) error {
	product := new(model.Product)
	db := database.NewPostgres()
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": err.Error()})
	}
	result := db.Db.Create(&product)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]string{"message": "product created"})
}
func GetAllProducts(c *fiber.Ctx) error {
	db := database.NewPostgres()
	var products []model.Product

	if err := db.Db.Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{"products": products})
}
