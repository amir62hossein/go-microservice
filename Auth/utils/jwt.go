package utils

import (
	"auth-service/model"
	"fmt"
	"time"

	"os"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateToken(user *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = fmt.Sprintf("%v", user.ID)       // Ensure user.ID is a string
	claims["email"] = fmt.Sprintf("%v", user.Email) // Ensure user.Email is a string
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(os.Getenv("KEY")))
	return t, err
}
