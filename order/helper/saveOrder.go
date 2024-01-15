package helper

import (
	"encoding/json"
	"fmt"
	"order-service/database"
	"order-service/model"
)

func SaveOrder(orderData []byte) ([]byte, error) {
	var order model.Order
	db := database.NewPostgres()
	err := json.Unmarshal(orderData, &order)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}
	if err := db.Db.Create(&order).Error; err != nil {
		return nil, fmt.Errorf("failed to save order to the database: %v", err)
	}

	fmt.Println("Order Saved:", order)

	productJson, err := json.Marshal(&order)
	if err != nil {
		panic(err)
	}
	return productJson, nil
}
