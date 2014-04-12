package controller

import (
	"fmt"
	"github.com/evantbyrne/evanbyrne-go/model/dto"
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"html/template"
	"net/http"
	"strings"
)

func GetView(request *http.Request, response http.ResponseWriter) (int, string) {
	url := util.TrimUrl(request)
	var db = new(util.Database)
	defer db.Close()
	if post, success := service.GetPostByUrl(db, url); success {
		layout, template, params := GetViewParams(post)
		return util.RespondTemplate(http.StatusOK, layout, template, params)
	}

	return PageNotFound(request, response)
}

func GetViewParams(post dto.Post) (string, string, map[string]interface{}) {
	params := make(map[string]interface{})
	params["layout"] = "default"
	params["design"] = "default"
	for _, meta := range post.Meta {
		if meta.Key == "content" || strings.HasPrefix(meta.Key, "content_") {
			params[meta.Key] = template.HTML(meta.Value)
		} else {
			params[meta.Key] = meta.Value
		}
	}

	layout := fmt.Sprintf("template/layout/%s.html", params["layout"])
	template := fmt.Sprintf("template/%s.html", params["design"])
	return layout, template, params
}