package util

import (
	"html/template"
	"io"
	"net/http"
)

func Respond(response http.ResponseWriter, status int, html string) {
	response.WriteHeader(status)
	io.WriteString(response, html)
	response.Header().Set("Content-Type", "text/html")
	response.Header().Set("Content-Length", string(len(html)))
}

func RespondTemplate(response http.ResponseWriter, status int, template_file string, data interface{}) {
	response.WriteHeader(status)
	t, _ := template.ParseFiles(template_file)
	t.Execute(response, data)
}