package api

import (
	"net/http"

	"github.com/maxwell7774/budgeting-backend/internal/database"
)

type ApiConfig struct {
	Resp http.ResponseWriter
	Req  *http.Request
	DB   *database.Queries
}
