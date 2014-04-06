package main

import (
	"github.com/evantbyrne/evanbyrne-go/model/service"
	"github.com/evantbyrne/evanbyrne-go/util"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		panic("Incorrect number of arguments")
	}

	var db = new(util.Database)
	defer db.Close()
	service.CreateUser(db, os.Args[1], os.Args[2])
}