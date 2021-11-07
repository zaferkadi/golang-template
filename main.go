package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/template-go-server/api"
	db "github.com/template-go-server/db/sqlc"
	"github.com/template-go-server/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	store := db.NewStore(conn)

	server, _ := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}
