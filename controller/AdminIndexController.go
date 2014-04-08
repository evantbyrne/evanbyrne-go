package controller

import (
	"github.com/evantbyrne/evanbyrne-go/model/dto"
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"net/http"
)

func GetAdminIndex(request *http.Request, response http.ResponseWriter) (int, string) {
	var db = new(util.Database)
	defer db.Close()
	if !service.ValidUserSession(db, request) {
		return util.Redirect(request, response, "/admin/login")
	}

	params := make(map[string][]dto.Post)
	params["posts"] = service.GetPostListing(db)
	return util.RespondTemplate(http.StatusOK, "template/layout/admin.html", "template/admin/index.html", params)
}