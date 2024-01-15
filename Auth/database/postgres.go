package database

import (
	"auth-service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Db *gorm.DB
}

func NewDB() (*Postgres, error) {
	dsn := "host=localhost user=postgres password=nourian1999 dbname=auth-service port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}
	return &Postgres{Db: db}, nil
}
