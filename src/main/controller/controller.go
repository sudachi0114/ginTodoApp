package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/sudachi0114/ginTodoApp/src/main/models"
)

// レコードを追加 (INSERT / CREATE)
func Insert(title string, description string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Error occured! Something wrong with insert...")
	}

	db.Create(&models.Todo{Title: title, Description: description, Done: "0"})

	defer db.Close()
}

// レコードの更新 (UPDATE)
func Update(id int, title string, description string, done string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Error occured! Something wrong with update...")
	}

	var todo models.Todo
	db.First(&todo, id)
	todo.Title = title
	todo.Description = description
	todo.Done = done
	db.Save(&todo)

	db.Close()
}

// レコードの削除 (DELETE)  ** FIXME: 物理削除 => 論理削除 **
func Delete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Error occured! Something wrong with delete...")
	}

	var todo models.Todo
	db.First(&todo, id)
	db.Delete(&todo)

	db.Close()
}

// Database のデータを全取得 (SELECT *)
func GetAll() []models.Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Error occured! Something wrong with fetch all data from database...")
	}

	var todos []models.Todo
	db.Order("created_at desc").Find(&todos)
	db.Find(&todos)

	db.Close()

	return todos
}

// Database のデータ 1件取得
func DbGetOne(id int) models.Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Error occured! Something wrong with fetch a data from database...")
	}

	var todo models.Todo
	db.First(&todo, id)

	db.Close()

	return todo
}
