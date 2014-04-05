package util

import (
	"bytes"
	"html/template"
)

func RespondTemplate(status int, template_file string, data interface{}) (int, string) {
	var buf bytes.Buffer 
	t, _ := template.ParseFiles(template_file)
	if err := t.Execute(&buf, data); err != nil {
		panic(err)
	}
	
	return status, buf.String()
}