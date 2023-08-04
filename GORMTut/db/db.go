package db

import (
	"log"

	db_types "github.com/ommahale/gormtut/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=root password=root dbname=Users port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func InitModels(db *gorm.DB) {
	db.AutoMigrate(&db_types.User{})
}
