package main

import (
	"fmt"
	"github.com/evantbyrne/evanbyrne-go/config"
	"github.com/evantbyrne/evanbyrne-go/model/dto"
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"html/template"
	"net/http"
	"strings"
)

func Router(response http.ResponseWriter, request *http.Request) {
	url := strings.TrimRight(request.URL.Path, "/")
	if url == "" {
		url = "/"
	}

	PageView(response, url)
}

func PageAdminIndex(response http.ResponseWriter, request *http.Request) {
	posts := service.GetPostListing()
	util.RespondTemplate(response, http.StatusOK, "template/admin/index.html", posts)
}

func PageAdminCreate(response http.ResponseWriter, request *http.Request) {
	params := make(map[string]string)
	if(request.Method == "POST") {
		request.ParseForm()
		content := request.PostForm.Get("content")
		data := util.Markdown(content)
		if url, ok := data["url"]; !ok || url == "" {
			params["content"] = content
			params["error"] = "Input error - URL not specified."
			util.RespondTemplate(response, http.StatusOK, "template/admin/create.html", params)
		} else {
			post := dto.Post{ Url: string(data["url"]) }
			for key, value := range data {
				if key != "url" {
					post.Meta = append(post.Meta, dto.PostMeta{ Key: key, Value: value })
				}
			}

			if err := service.CreatePost(post); err != nil {
				params["content"] = content
				params["error"] = "Internal server error - " + err.Error()
				util.RespondTemplate(response, http.StatusInternalServerError, "template/admin/create.html", params)
			} else {
				http.Redirect(response, request, "/admin", 303)
			}
		}
	} else {
		util.RespondTemplate(response, http.StatusOK, "template/admin/create.html", params)
	}
}

func PageView(response http.ResponseWriter, url string) {
	params := make(map[string]interface{})
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
		util.RespondTemplate(response, http.StatusOK, template, params)
	} else {
		params["url"] = url
		util.RespondTemplate(response, http.StatusNotFound, "template/404.html", params)
	}
}

func main() {
	// /admin/create -> PageAdminCreate
	// /admin/delete/{id:[0-9]+} -> PageAdminDelete
	// /admin/edit/{id:[0-9]+} -> PageAdminEdit
	// /admin/login -> PageAdminLogin
	// * -> PageView
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/admin", PageAdminIndex)
	http.HandleFunc("/admin/create", PageAdminCreate)
	http.HandleFunc("/", Router)
	http.ListenAndServe(":" + config.HttpPort, nil)
}