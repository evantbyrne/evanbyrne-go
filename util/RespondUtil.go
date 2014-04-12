package util

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"
)

func Redirect(request *http.Request, response http.ResponseWriter, url string) (int, string) {
	http.Redirect(response, request, url, 303)
	return 303, ""
}

func RespondTemplate(status int, templateLayoutFile string, templateFile string, data interface{}) (int, string) {
	var buf bytes.Buffer 
	t, _ := template.ParseFiles(templateLayoutFile, templateFile)
	if err := t.Execute(&buf, data); err != nil {
		panic(err)
	}
	
	return status, buf.String()
}

func TrimUrl(request *http.Request) string {
	url := strings.TrimRight(request.URL.Path, "/")
	if url == "" {
		return "/"
	}
	
	return url
}