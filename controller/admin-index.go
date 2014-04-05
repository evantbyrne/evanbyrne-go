package controller

import (
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"net/http"
)

func GetAdminIndex() (int, string) {
	posts := service.GetPostListing()
	return util.RespondTemplate(http.StatusOK, "template/admin/index.html", posts)
}