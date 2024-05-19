package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open("host=exchange_db user=postgres dbname=my_database password=password port=5433 sslmode=disable"), &gorm.Config{},
	)
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	return db
}
