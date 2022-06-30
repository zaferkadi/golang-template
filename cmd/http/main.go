package main

import (
	"database/sql"
	"io/ioutil"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/template-go-server/api"
	db "github.com/template-go-server/db/sqlc"
	"github.com/template-go-server/util"
)

// list all directories in current directory
func listAllDirs() {
	log.Println("List all directories in current directory")
	dirs, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal("Cannot read current directory:", err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			log.Println(dir.Name())
		}
	}
}

func runMigrations(config util.Config) {

	log.Println("DBDriver:", config.DBDriver)
	log.Println("DBSource:", config.DBSource)

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Panic("Migrate: Open Error ", err)
	}
	defer db.Close()
	// why connection refused?
	db.Ping()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Panic("Migrate: WithInstance Error ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"simple_books", driver)
	if err != nil {
		log.Panic("Migrate: NewWithDatabaseInstance Error ", err)
	}
	err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	if err != nil && err != migrate.ErrNoChange {
		log.Panic("up issue ", err)
	}

}

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	// runMigrations
	runMigrations(config)

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}
