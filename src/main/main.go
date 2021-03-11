package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"github.com/sudachi0114/ginTodoApp/src/main/controller"
)

func main() {
	fmt.Printf("< Server started >\n")

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	// index root
	router.GET("/", func(ctx *gin.Context) {

		todos := controller.GetAll()

		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	// Create Todo
	router.POST("/new", func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		description := ctx.PostForm("description")
		controller.Insert(title, description)

		ctx.Redirect(302, "/")
	})

	// Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := controller.DbGetOne(id)

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
		var done string = "0"
		if ctx.PostForm("done") == "1" {
			done = "1"
		}

		controller.Update(id, title, description, done)

		ctx.Redirect(302, "/")
	})

	router.Run()
}
