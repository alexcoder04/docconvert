package main

import (
	"embed"
	"flag"
	"fmt"
	"text/template"

	"github.com/alexcoder04/docconvert/handlers"
	"github.com/gin-gonic/gin"
)

//go:embed templates/index.html
var TemplateFiles embed.FS

var (
	Template *template.Template

	ApiOnly  = flag.Bool("api-only", false, "run the program in API-only mode")
	AllHosts = flag.Bool("all-hosts", false, "listen on 0.0.0.0")
	Port     = flag.Int("port", 8080, "port number to listen on")
)

func LoadTemplate() (*template.Template, error) {
	tmpl := template.New("")

	contents, err := TemplateFiles.ReadFile("templates/index.html")
	if err != nil {
		return nil, err
	}

	_, err = tmpl.Parse(string(contents))
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func main() {
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	if !*ApiOnly {
		var err error
		Template, err = LoadTemplate()
		if err != nil {
			panic("failed to load html template")
		}

		router.GET("/", func(c *gin.Context) {
			Template.Execute(c.Writer, gin.H{})
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
