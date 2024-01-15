package database

import (
	"gorm.io/gorm"
	"order-service/model"

	"gorm.io/driver/postgres"
)

type Postgres struct {
	Db *gorm.DB
}

func NewPostgres() *Postgres {

	dsn := "host=localhost user=postgres password=nourian1999 dbname=order-service port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.Order{})
	if err != nil {
		panic(err.Error())
	}
	return &Postgres{Db: db}
}
