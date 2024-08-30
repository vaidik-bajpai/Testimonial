package main

import (
	"context"
	"flag"

	"github.com/vaidik-bajpai/testimonials/db"
	"github.com/vaidik-bajpai/testimonials/handler"
	"github.com/vaidik-bajpai/testimonials/storer"
	"github.com/vaidik-bajpai/testimonials/validate"
)

func main() {
	var (
		dsn = flag.String("DB-DSN", "mongodb://localhost:27017", "uri to database that tell the api which db to connect to")
	)

	ok := validate.RegisterValidators()
	if !ok {
		panic("error registering validators")
	}

	database, err := db.NewDatabase(*dsn)
	if err != nil {
		panic("error connecting to the database")
	}
	defer database.Close()

	client := database.GetDB()

	spaceCollection := storer.MakeCollection(client, "Space")

	st := storer.NewStorer(client, &storer.Collections{Space: spaceCollection})
	hdl := handler.NewHandler(context.Background(), st)
	handler.RegisterRoutes(hdl)
	handler.Start(":8080")
}
