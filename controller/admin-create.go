package controller

import (
	"github.com/evantbyrne/evanbyrne-go/model/dto"
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"net/http"
)

func GetAdminCreate() (int, string) {
	params := make(map[string]string)
	return util.RespondTemplate(http.StatusOK, "template/admin/create.html", params)
}

func PostAdminCreate(request *http.Request, response http.ResponseWriter) (int, string) {
	params := make(map[string]string)
	request.ParseForm()
	content := request.PostForm.Get("content")
	data := util.Markdown(content)
	if url, ok := data["url"]; !ok || url == "" {
		params["content"] = content
		params["error"] = "Input error - URL not specified."
		return util.RespondTemplate(http.StatusOK, "template/admin/create.html", params)
	}

	post := dto.Post{ Url: string(data["url"]) }
	for key, value := range data {
		if key != "url" {
			post.Meta = append(post.Meta, dto.PostMeta{ Key: key, Value: value })
		}
	}

	if err := service.CreatePost(post); err != nil {
		params["content"] = content
		params["error"] = "Internal server error - " + err.Error()
		return util.RespondTemplate(http.StatusInternalServerError, "template/admin/create.html", params)
	}

	http.Redirect(response, request, "/admin", 303)
	return 303, ""
}