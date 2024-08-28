package main

import (
	"context"
	"flag"

	"github.com/vaidik-bajpai/testimonials/db"
	"github.com/vaidik-bajpai/testimonials/handler"
	"github.com/vaidik-bajpai/testimonials/storer"
)

func main() {
	var (
		dsn = flag.String("DB-DSN", "mongodb://localhost:27017", "uri to database that tell the api which db to connect to")
	)

	database, err := db.NewDatabase(*dsn)
	if err != nil {
		panic("error connecting to the database")
	}
	defer database.Close()

	st := storer.NewStorer(database.GetDB())
	hdl := handler.NewHandler(context.Background(), st)
	handler.RegisterRoutes(hdl)
	handler.Start(":8080")
}
