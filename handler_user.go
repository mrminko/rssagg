package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mrminko/rssagg/internal/database"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerUserGet(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, dbUserToUser(user))
}

func (apiCfg *apiConfig) handlerUserCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error when parsing JSON:%v", err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot create user: %v", err))
		return
	}

	respondWithJSON(w, 201, dbUserToUser(user))
}
