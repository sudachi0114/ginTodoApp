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
	Text string
	Status string
}

// Database 初期化 (migration)
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開なかったわ!ごめん!")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

// Database にレコードを追加
func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開なかったわ!ごめん!")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

// Database にあるレコードの更新
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開なかったわ!ごめん!")
	}

	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)

	db.Close()
}

// Database にあるレコードの削除
func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開なかったわ!ごめん!")
	}

	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)

	db.Close()
}

// Database のデータを全取得
func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開なかったわ!ごめん!")
	}

	var todos []Todo
	// db.Order("create_at desc").Find(&todos)
	db.Find(&todos)

	db.Close()

	return todos
}

// Database のデータ 1件取得
func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開なかったわ!ごめん!")
	}

	var todo Todo
	db.First(&todo, id)

	db.Close()
	return todo
}



func main() {
	fmt.Printf("< Program started >\n")

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	fmt.Println("<info> DataBase Inited!!")
	dbInit()

	x := "すだちです"
	var todo Todo
	todo.Text = "このアプリを仕上げる"
	todo.Status = "yet"
	// todos := [1] Todo{todo}
	fmt.Println(todo.Text, todo.Status)

	// index root
	router.GET("/", func(ctx *gin.Context) {

		todos := dbGetAll()
		
		ctx.HTML(200, "index.html", gin.H{
			"data": x,
			"todo": todo,
			"todos": todos,
		})
	})

	// Create Todo
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbInsert(text, status)

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

	router.Run()
}
