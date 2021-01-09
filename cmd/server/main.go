package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"weight-tracker/pkg/api"
	"weight-tracker/pkg/app"
	"weight-tracker/pkg/repository"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

// func run will be responsible for setting up db connections, routers etc
func run() error {
	// I'm used to working with postgres feel free to use any db you would like
	// I'm not going to cover how to create a database here but create a database
	// and call it something along the lines of "weight tracker"
	connectionString := "postgres://postgres:postgres@localhost/weight_tracker?sslmode=disable"

	// setup database connection
	db, err := setupDatabase(connectionString)
	if err != nil {
		return err
	}

	// create storage dependency
	storage := repository.NewStorage(db)

	// create router dependecy
	router := gin.Default()
	router.Use(cors.Default())

	// create user service
	userService := api.NewUserService(storage)

	// create weight service
	weightService := api.NewWeightService(storage)

	server := app.NewServer(router, userService, weightService)

	// start the server
	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}

func setupDatabase(connString string) (*sql.DB, error) {
	// change "postgres" for whatever supported database you want to use
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	// ping the DB to ensure that it is connected
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
