package main

import (
	"flag"
	"net/http"

	"github.com/alexcoder04/docconvert/handlers"
	"github.com/gin-gonic/gin"
)

var ApiOnly = flag.Bool("api-only", false, "run the program in API-only mode")

func main() {
	flag.Parse()

	router := gin.Default()

	if !*ApiOnly {
		router.LoadHTMLGlob("templates/*.html")

		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
	}

	router.POST("/convert/upload", handlers.HandleFileUpload)
	router.POST("/convert/text", handlers.HandleTextData)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
