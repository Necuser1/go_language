package comman

import (
	"gorm.io/gorm"
)

type Book struct {
	Id        string `gorm:"primary key;autoIncrement" json:"id"`
	Name      string `json:"name"`
	Auther    string `json:"auther"`
	Publisher string `json:"publisher"`
	Edision   string `json:"edision"`
}

func MigrateBook(db *gorm.DB) error {
	err := db.AutoMigrate(&Book{})
	return err
}
