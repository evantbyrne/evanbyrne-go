package controller

import (
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"net/http"
)

func GetAdminLogin() (int, string) {
	params := make(map[string]string)
	params["title"] = "Log In"
	return util.RespondTemplate(http.StatusOK, "template/layout/admin.html", "template/admin/login.html", params)
}

func PostAdminLogin(request *http.Request, response http.ResponseWriter) (int, string) {
	request.ParseForm()
	email := request.PostForm.Get("email")
	password := request.PostForm.Get("password")
	if email != "" && password != "" {
		var db = new(util.Database)
		defer db.Close()
		if service.LoginUser(db, response, email, password) {
			return util.Redirect(request, response, "/admin")
		}
	}

	params := make(map[string]string)
	params["title"] = "Log In"
	params["email"] = email
	params["error"] = "Input error - Invalid login credentials."
	return util.RespondTemplate(http.StatusOK, "template/layout/admin.html", "template/admin/login.html", params)
}