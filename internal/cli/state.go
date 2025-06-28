package cli

import "github.com/maxwell7774/budgeting-backend/internal/database"

type State struct {
	db *database.Queries
}

func NewState(db *database.Queries) *State {
	return &State{
		db: db,
	}
}
