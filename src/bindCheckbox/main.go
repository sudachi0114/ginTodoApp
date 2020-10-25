package main

import (
	"github.com/gin-gonic/gin"
)

type myForm struct {
	Status string `form:"mycheckbox"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	r.GET("/", indexHandler)
	r.POST("/", formHandler)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello"})
	})
	r.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "form.html", gin.H{})
}

func formHandler(c *gin.Context) {
	var myform myForm
	c.Bind(&myform)
	c.JSON(200, gin.H{"mycheckbox": myform.Status})
}

// cf.
//  https://github.com/gin-gonic/gin/issues/129#issuecomment-124260092
//  https://chenyitian.gitbooks.io/gin-web-framework/content/docs/21.html
