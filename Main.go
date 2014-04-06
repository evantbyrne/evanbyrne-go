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
	m.Get("/admin/edit/:id", controller.GetAdminEdit)
	m.Post("/admin/edit", controller.PostAdminEdit)
	m.Get("/admin/login", controller.GetAdminLogin)
	m.Post("/admin/login", controller.PostAdminLogin)
	m.Get("**", controller.GetView)
	http.ListenAndServe(":" + config.HttpPort, m)
}