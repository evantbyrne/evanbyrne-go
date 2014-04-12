package controller

import (
	"fmt"
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"html/template"
	"net/http"
)

func GetView(request *http.Request, response http.ResponseWriter) (int, string) {
	params := make(map[string]interface{})
	url := util.TrimUrl(request)
	var db = new(util.Database)
	defer db.Close()
	if post, success := service.GetPostByUrl(db, url); success {
		params["layout"] = "default"
		params["design"] = "default"
		for _, meta := range post.Meta {
			if meta.Key == "content" {
				params["content"] = template.HTML(meta.Value)
			} else {
				params[meta.Key] = meta.Value
			}
		}

		layout := fmt.Sprintf("template/layout/%s.html", params["layout"])
		template := fmt.Sprintf("template/%s.html", params["design"])
		return util.RespondTemplate(http.StatusOK, layout, template, params)
	}

	return PageNotFound(request, response)
}