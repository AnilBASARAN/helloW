package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	basePath := filepath.Join("templates", "base.layout.tmpl")
	pagePath := filepath.Join("templates", tmpl)

	parsedTemplate, err := template.ParseFiles(basePath, pagePath)
	if err != nil {
		fmt.Println("error parsing templates:", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = parsedTemplate.ExecuteTemplate(w, "base", nil)
	if err != nil {
		fmt.Println("error executing template:", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

//lets consider an alternative approach
