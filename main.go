package main

import (
	"flag"
	"fmt"

	"github.com/alexcoder04/docconvert/handlers"
	"github.com/gin-gonic/gin"
)

var (
	AllHosts = flag.Bool("all-hosts", false, "listen on 0.0.0.0")
	Port     = flag.Int("port", 8080, "port number to listen on")
	Debug    = flag.Bool("debug", false, "run in debug mode")
)

func main() {
	flag.Parse()

	if !*Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	router.StaticFS("/static", handlers.StaticFS)

	router.GET("/", handlers.HandleIndex)
	router.POST("/img2text", handlers.HandleImg2Text)

	router.POST("/convert/upload", handlers.HandleFileUpload)
	router.POST("/convert/text", handlers.HandleTextData)

	if *AllHosts {
		router.Run(fmt.Sprintf(":%d", *Port))
	} else {
		router.Run(fmt.Sprintf("127.0.0.1:%d", *Port))
	}
}
