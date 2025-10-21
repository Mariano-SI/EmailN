package database

import (
	"emailn/internal/domain/campaign"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=emailn password=d4#rt6 dbname=emailn port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database. Please check your connection settings.")
	}

	log.Println("✅ Successfully connected to the database.")
	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})
	return db
}
