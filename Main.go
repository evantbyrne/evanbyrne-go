package main

import (
	"github.com/evantbyrne/evanbyrne-go/config"
	"github.com/evantbyrne/evanbyrne-go/controller"
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Get("/static/**", http.StripPrefix("/static", http.FileServer(http.Dir(config.Static))))
	m.Get("/admin", controller.GetAdminIndex)
	m.Get("/admin/create", controller.GetAdminCreate)
	m.Post("/admin/create", controller.PostAdminCreate)
	m.Get("/admin/delete/:id", controller.GetAdminDelete)
	m.Post("/admin/delete", controller.PostAdminDelete)
	m.Get("/admin/edit/:id", controller.GetAdminEdit)
	m.Post("/admin/edit", controller.PostAdminEdit)
	m.Get("/admin/login", controller.GetAdminLogin)
	m.Post("/admin/login", controller.PostAdminLogin)
	m.Post("/admin/logout", controller.PostAdminLogout)
	m.Get("/blog", controller.GetBlog)
	m.Get("/", controller.GetHome)
	m.Get("**", controller.GetView)
	m.Post("**", controller.PageNotFound)
	http.ListenAndServe(":" + config.HttpPort, m)
}