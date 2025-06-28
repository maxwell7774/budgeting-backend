package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/maxwell7774/budgeting-backend/internal/database"
)

func HandlerUserCreate(cfg *ApiConfig) {
	type parameters struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	}

	decoder := json.NewDecoder(cfg.Req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(cfg.Resp, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.DB.CreateUser(cfg.Req.Context(), database.CreateUserParams{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
	})
	if err != nil {
		respondWithError(cfg.Resp, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}

	type response struct {
		ID        string    `json:"id"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	respondWithJSON(cfg.Resp, http.StatusCreated, response{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}
