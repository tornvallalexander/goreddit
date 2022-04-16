package main

import (
	"database/sql"
	"github.com/tornvallalexander/goreddit/api"
	db "github.com/tornvallalexander/goreddit/db/sqlc"
	"github.com/tornvallalexander/goreddit/utils"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("unable to load environment variables:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
