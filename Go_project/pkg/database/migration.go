package database

import (
	"GO_PROJECT/pkg/comman"

	"gorm.io/gorm"
)

func MigrateBook(db *gorm.DB) error {
	err := db.AutoMigrate(&comman.Book{})
	return err
}

func Migrate(db *gorm.DB) error {
	err := MigrateBook(db)
	if err != nil {
		return err
	}
	return nil
}
