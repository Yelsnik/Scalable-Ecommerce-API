package main

import (
	db "auction-service/db/sqlc"
	"auction-service/util"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// connect to database
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("could not connect", err)
	}

	_ = db.NewStore(conn)
}
