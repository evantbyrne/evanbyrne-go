package util

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/evantbyrne/evanbyrne-go/config"
	"github.com/evantbyrne/evanbyrne-go/model/dto"
	"log"
	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
	Opened bool
}

func (db *Database) Open() *sql.DB {
	if !db.Opened {
		connection, err := sql.Open("postgres", config.Database)
		if err != nil {
			log.Fatal(err)
		}

		db.Connection = connection
		db.Opened = true
	}

	return db.Connection
}

func (db *Database) Gorp() *gorp.DbMap {
	connection := db.Open()
	dbmap := &gorp.DbMap{ Db: connection, Dialect: gorp.PostgresDialect{} }
	dbmap.AddTableWithName(dto.Post{}, "post").SetKeys(true, "Id")
	dbmap.AddTableWithName(dto.PostMeta{}, "post_meta")
	return dbmap
}

func (db *Database) Close() {
	if db.Opened {
		db.Opened = false
		db.Connection.Close()
	}
}