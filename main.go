package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/alexcoder04/docconvert/handlers"
	"github.com/gin-gonic/gin"
)

var (
	ApiOnly = flag.Bool("api-only", false, "run the program in API-only mode")
	Port    = flag.Int("port", 8080, "port number to listen on")
)

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

	router.Run(fmt.Sprintf(":%d", *Port))
}
