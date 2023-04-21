package handlers

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed templates
var TemplateFiles embed.FS

//go:embed static
var StaticFiles embed.FS

var (
	Templates *template.Template
	StaticFS  http.FileSystem
)

func init() {
	// load templates
	var err error
	Templates, err = template.ParseFS(TemplateFiles, "templates/*.html")
	if err != nil {
		panic(fmt.Sprintf("failed to load html templates, %s", err.Error()))
	}

	// strip prefix from static fs
	sub, err := fs.Sub(StaticFiles, "static")
	if err != nil {
		panic("failed to strip 'static' from static files")
	}
	StaticFS = http.FS(sub)
}
