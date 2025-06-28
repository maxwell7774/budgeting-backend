package app

import (
	"log"
	"net/http"

	"github.com/maxwell7774/budgeting-backend/internal/api"
	"github.com/maxwell7774/budgeting-backend/internal/database"
)

type App struct {
	db   *database.Queries
	mux  *http.ServeMux
	port string
}

type AppHandler func(*api.ApiConfig)

func NewApp(port string, db *database.Queries) *App {
	return &App{
		db:   db,
		mux:  http.NewServeMux(),
		port: port,
	}
}

func (app *App) HandleFunc(route string, handler AppHandler) {
	app.mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		c := api.ApiConfig{
			Resp:  w,
			Req: r,
			DB:      app.db,
		}
		handler(&c)
	})
}

func (app *App) Start() {
	app.HandleFunc("GET /api/v1/hello", api.HandlerHello)
	app.HandleFunc("POST /api/v1/users", api.HandlerUserCreate)

	log.Printf("Listening on port %s", app.port)
	log.Fatal(http.ListenAndServe(app.port, app.mux))
}
