package model

type Order struct {
	UserId      string `json:"user_id" gorm:"primaryKey"`
	Email       string `json:"email"`
	ProductName string `json:"productName"`
}
