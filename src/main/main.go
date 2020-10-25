package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// data-object というやつ??
type Todo struct {
	gorm.Model
	Title string
	Description string
	Done bool
}

// Database 初期化 (migration)
func migrateDB() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Error occured! Something wrong with migration...")
	}
	db.AutoMigrate(&Todo{})

	defer db.Close()
}

// レコードを追加 (INSERT / CREATE)
func insert(title string, description string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Error occured! Something wrong with insert...")
	}

	db.Create(&Todo{Title: title, Description: description, Done: false})

	defer db.Close()
}

// レコードの更新 (UPDATE)
func update(id int, title string, description string, done bool) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Error occured! Something wrong with update...")
	}

	var todo Todo
	db.First(&todo, id)
	todo.Title = title
	todo.Description = description
	todo.Done = done
	db.Save(&todo)

	db.Close()
}

// レコードの削除 (DELETE)  ** FIXME: 物理削除 => 論理削除 **
func delete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Error occured! Something wrong with delete...")
	}

	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)

	db.Close()
}

// Database のデータを全取得 (SELECT *)
func getAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Error occured! Something wrong with fetch all data from database...")
	}

	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Find(&todos)

	db.Close()

	return todos
}

// Database のデータ 1件取得
func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Error occured! Something wrong with fetch a data from database...")
	}

	var todo Todo
	db.First(&todo, id)

	db.Close()

	return todo
}



func main() {
	fmt.Printf("< Server started >\n")

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	fmt.Println("<info> DataBase Migration!!")
	migrateDB()

	// index root
	router.GET("/", func(ctx *gin.Context) {

		todos := getAll()
		
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	// Create Todo
	router.POST("/new", func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		description := ctx.PostForm("description")
		insert(title, description)

		ctx.Redirect(302, "/")
	})

	// Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dbGetOne(id)

		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		title := ctx.PostForm("title")
		description := ctx.PostForm("description")
		var done bool = false
		if ctx.PostForm("done") == "1" {
			done = true
		}

		update(id, title, description, done)

		ctx.Redirect(302, "/")
	})

	router.Run()
}
