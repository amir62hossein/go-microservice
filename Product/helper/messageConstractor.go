package helper

import "encoding/json"

func MessageConstructor(id, email, productName string) []byte {
	productMessage := struct {
		UserId      string `json:"user_id"`
		Email       string `json:"email"`
		ProductName string `json:"productName"`
	}{
		UserId:      id,
		Email:       email,
		ProductName: productName,
	}
	jsonMessage, err := json.Marshal(&productMessage)
	if err != nil {
		panic(err.Error())
	}
	return jsonMessage
}
