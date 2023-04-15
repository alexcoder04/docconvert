package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/alexcoder04/docconvert/handlers"
	"github.com/gin-gonic/gin"
)

var (
	ApiOnly  = flag.Bool("api-only", false, "run the program in API-only mode")
	AllHosts = flag.Bool("all-hosts", false, "listen on 0.0.0.0")
	Port     = flag.Int("port", 8080, "port number to listen on")
)

func main() {
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	if !*ApiOnly {
		router.LoadHTMLGlob("templates/*.html")

		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
	}

	router.POST("/convert/upload", handlers.HandleFileUpload)
	router.POST("/convert/text", handlers.HandleTextData)

	if *AllHosts {
		router.Run(fmt.Sprintf(":%d", *Port))
	} else {
		router.Run(fmt.Sprintf("127.0.0.1:%d", *Port))
	}
}
