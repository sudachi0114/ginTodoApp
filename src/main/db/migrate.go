package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"github.com/sudachi0114/ginTodoApp/src/main/models"
)

// Database migration
func main() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Error occured! Something wrong with migration...")
	}
	db.AutoMigrate(&models.Todo{})

	defer db.Close()
}
