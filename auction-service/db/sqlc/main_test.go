package db

import (
	"auction-service/util"
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// connect to db
	testDB, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("could not connect", err)
	}

	testStore = NewStore(testDB)

	os.Exit(m.Run())
}
