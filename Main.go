package main

import (
	"fmt"
	"github.com/evantbyrne/evanbyrne-go/config"
	"github.com/evantbyrne/evanbyrne-go/model/dto"
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"github.com/go-martini/martini"
	"html/template"
	"net/http"
	"strings"
)

func GetAdminIndex() (int, string) {
	posts := service.GetPostListing()
	return util.RespondTemplate(http.StatusOK, "template/admin/index.html", posts)
}

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

func main() {
	m := martini.Classic()
	m.Get("/static/**", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	m.Get("/admin", GetAdminIndex)
	m.Get("/admin/create", GetAdminCreate)
	m.Post("/admin/create", PostAdminCreate)
	m.Get("**", GetView)
	http.ListenAndServe(":" + config.HttpPort, m)
}