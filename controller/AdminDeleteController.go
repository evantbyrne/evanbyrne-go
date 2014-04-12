package controller

import (
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"github.com/go-martini/martini"
	"net/http"
)

func GetAdminDelete(request *http.Request, response http.ResponseWriter, uri martini.Params) (int, string) {
	var db = new(util.Database)
	defer db.Close()
	if !service.ValidUserSession(db, request) {
		return util.Redirect(request, response, "/admin/login")
	}

	params := make(map[string]string)
	if post, success := service.GetPostById(db, uri["id"]); success {
		params["id"] = uri["id"]
		params["url"] = post.Url
		params["title"] = "Delete"
		return util.RespondTemplate(http.StatusOK, "template/layout/admin.html", "template/admin/delete.html", params)
	}

	return PageNotFound(request, response)
}

func PostAdminDelete(request *http.Request, response http.ResponseWriter) (int, string) {
	var db = new(util.Database)
	defer db.Close()
	if !service.ValidUserSession(db, request) {
		return util.Redirect(request, response, "/admin/login")
	}

	request.ParseForm()
	id := request.PostForm.Get("id")
	if _, success := service.GetPostById(db, id); success {
		service.DeletePostById(db, id)
		return util.Redirect(request, response, "/admin")
	}
	
	return PageNotFound(request, response)
}