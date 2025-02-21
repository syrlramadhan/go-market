package controller

import (
	"html/template"
	"net/http"
)

func RenderTemplate(writer http.ResponseWriter,tmpl string, data map[string]interface{}){
	layoutPath := "templates/layout.html"
	pagePath := "templates/pages/"+tmpl

	t, err := template.ParseFiles(layoutPath, pagePath)
	if err != nil {
		http.Error(writer, "Template rendering error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(writer, data)
	if err != nil {
		http.Error(writer, "Failed to execute template", http.StatusInternalServerError)
		return
	}
}