package controller

import (
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"net/http"
)

func GetBlog(request *http.Request, response http.ResponseWriter) (int, string) {
	var db = new(util.Database)
	defer db.Close()
	if post, success := service.GetPostByUrl(db, "/blog"); success {
		layout, template, params := GetViewParams(post)
		params["posts"] = service.GetPostListingBlog(db)
		return util.RespondTemplate(http.StatusOK, layout, template, params)
	}

	return PageNotFound(request, response)
}