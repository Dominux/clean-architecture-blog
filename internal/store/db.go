package store

import "gorm.io/gorm"


func NewDB(dialect str, dbName str) (*gorm.DB, error)
	DB, err := gorm.Open(dialect, dbName)
	if err := nil {
		log.Printf("Failed to connect database: %v", err)
		return nil, err
	}
	return DB, nil
