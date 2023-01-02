package database

import (
	"GO_PROJECT/pkg/comman"
)

//var dbstatus map[string]bool

func InsertData(data comman.Book) bool {
	db := DbConnection()
	result := db.Create(&data)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func DeleteBook(bookID string) bool {
	db := DbConnection()
	book := comman.Book{}
	result := db.Where("id = ?", bookID).Delete(&book)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
func GetAllBook() *[]comman.Book {
	db := DbConnection()
	book := &[]comman.Book{}
	db.Find(book)
	return book
}

func GetById(bookId string) comman.Book {
	db := DbConnection()
	book := comman.Book{}
	_ = db.Where("id = ?", bookId).Find(&book)
	return book
}
