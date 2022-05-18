package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/template-go-server/util"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannon load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
