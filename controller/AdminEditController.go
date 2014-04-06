package controller

import (
	"github.com/evantbyrne/evanbyrne-go/model/dto"
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"github.com/go-martini/martini"
	"net/http"
)

func GetAdminEdit(request *http.Request, response http.ResponseWriter, uri martini.Params) (int, string) {
	var db = new(util.Database)
	defer db.Close()
	if !service.ValidUserSession(db, request) {
		return util.Redirect(request, response, "/admin/login")
	}

	params := make(map[string]string)
	if post, success := service.GetPostById(db, uri["id"]); success {
		params["id"] = uri["id"]
		for _, meta := range post.Meta {
			if meta.Key == "markdown" {
				params["markdown"] = meta.Value
			}
		}

		return util.RespondTemplate(http.StatusOK, "template/admin/edit.html", params)
	}

	return util.RespondTemplate(http.StatusNotFound, "template/404.html", params)
}

func PostAdminEdit(request *http.Request, response http.ResponseWriter) (int, string) {
	var db = new(util.Database)
	defer db.Close()
	if !service.ValidUserSession(db, request) {
		return util.Redirect(request, response, "/admin/login")
	}

	params := make(map[string]string)
	request.ParseForm()
	id := request.PostForm.Get("id")
	if post, success := service.GetPostById(db, id); success {
		content := request.PostForm.Get("content")
		data := util.Markdown(content)
		if url, ok := data["url"]; !ok || url == "" {
			params["id"] = id
			params["markdown"] = content
			params["error"] = "Input error - URL not specified."
			return util.RespondTemplate(http.StatusOK, "template/admin/edit.html", params)
		}

		post.Url = data["url"]
		post.Meta = make([]dto.PostMeta, 1)
		for key, value := range data {
			if key != "url" {
				post.Meta = append(post.Meta, dto.PostMeta{ Key: key, Value: value })
			}
		}

		if err := service.EditPost(db, post); err != nil {
			params["id"] = id
			params["markdown"] = content
			params["error"] = "Internal server error - " + err.Error()
			return util.RespondTemplate(http.StatusInternalServerError, "template/admin/edit.html", params)
		}

		http.Redirect(response, request, "/admin", 303)
		return 303, ""
	}

	return util.RespondTemplate(http.StatusNotFound, "template/404.html", params)
}