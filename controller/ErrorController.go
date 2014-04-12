package controller

import (
	"github.com/evantbyrne/evanbyrne-go/util"
	"net/http"
	"strings"
)

func PageNotFound(request *http.Request, response http.ResponseWriter) (int, string) {
	params := make(map[string]string)
	params["title"] = "Page not found"
	params["url"] = util.TrimUrl(request)
	var layout string
	if strings.HasPrefix(params["url"], "/admin") {
		layout = "template/layout/admin.html"
	} else {
		layout = "template/layout/default.html"
	}

	return util.RespondTemplate(http.StatusNotFound, layout, "template/404.html", params)
}