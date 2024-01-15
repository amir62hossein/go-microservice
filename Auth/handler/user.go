package handler

import (
	"auth-service/database"
	"auth-service/model"
	"auth-service/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {

	user := new(model.User)
	foundedUser := new(model.User)
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": err.Error()})
	}

	db, _ := database.NewDB()

	if err := db.Db.Where("email = ?", user.Email).Find(&foundedUser).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": err.Error()})
	}

	if user.Password != foundedUser.Password {
		var err = errors.New("invalid Username Or Password")
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": err.Error()})
	}

	token, err := utils.GenerateToken(foundedUser)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": err.Error()})
	}

	return ctx.JSON(map[string]string{"token": token})
}
func Register(ctx *fiber.Ctx) error {
	user := new(model.User)
	db, _ := database.NewDB()
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]string{"message": err.Error()})
	}

	if err := db.Db.Create(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(map[string]string{"message": "User Created"})
}
