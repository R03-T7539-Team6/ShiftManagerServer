package main

import (
	"github.com/R03-T7539-Team6/ShiftManagerSerer/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	db.Init()
	defer db.Close()
	r.Run() // listen and serve on localhost:8080
}
