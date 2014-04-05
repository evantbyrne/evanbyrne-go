package controller

import (
	"fmt"
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"html/template"
	"net/http"
	"strings"
)

func GetView(request *http.Request) (int, string) {
	params := make(map[string]interface{})
	url := strings.TrimRight(request.URL.Path, "/")
	if url == "" {
		url = "/"
	}

	if post, success := service.GetPostByUrl(url); success {
		params["design"] = "default"
		for _, meta := range post.Meta {
			if meta.Key == "content" {
				params["content"] = template.HTML(meta.Value)
			} else {
				params[meta.Key] = meta.Value
			}
		}

		template := fmt.Sprintf("template/%s.html", params["design"])
		return util.RespondTemplate(http.StatusOK, template, params)
	}

	params["url"] = url
	return util.RespondTemplate(http.StatusNotFound, "template/404.html", params)
}