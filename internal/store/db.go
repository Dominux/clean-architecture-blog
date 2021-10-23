package store

import (
	"log"

	"gorm.io/gorm"
)

func NewDB(d gorm.Dialector) (*gorm.DB, error) {
	return NewDBWithConfig(d, &gorm.Config{})
}

func NewDBWithConfig(d gorm.Dialector, c *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(d, c)
	if err != nil {
		log.Printf("Failed to connect database: %v", err)
		return nil, err
	}
	return db, nil
}
