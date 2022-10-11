package db

import (
	"log"

	"github.com/adilsonmenechini/goapi/internal/entity"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return err
	}
	log.Println("migrated")
	return nil
}
