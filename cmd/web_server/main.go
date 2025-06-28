package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/maxwell7774/budgeting-backend/internal/app"
	"github.com/maxwell7774/budgeting-backend/internal/database"
)

func main() {
	port := ":8080"

	db, err := sql.Open("postgres", "postgres://stitch:@localhost:5432/budgeting?sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	dbQueries := database.New(db)

	app := app.NewApp(
		port,
		dbQueries,
	)

	app.Start()
}
