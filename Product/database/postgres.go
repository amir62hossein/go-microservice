package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"product-service/model"
)

type Postgres struct {
	Db *gorm.DB
}

func NewPostgres() *Postgres {
	dsn := "host=localhost user=postgres password=nourian1999 dbname=product-service port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&model.Product{})
	return &Postgres{Db: db}
}
