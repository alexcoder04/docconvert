package handlers

import (
	"embed"
	"fmt"
	"text/template"
)

//go:embed templates/index.html
var TemplateFiles embed.FS

var Template *template.Template

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

func Setup(webgui bool) {
	var err error
	Template, err = LoadTemplate()
	if err != nil {
		panic(fmt.Sprintf("failed to load html template, %s", err.Error()))
	}
}
