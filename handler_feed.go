package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mrminko/rssagg/internal/database"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerFeedGet(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, dbUserToUser(user))
}

func (apiCfg *apiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error when parsing JSON:%v", err))
		return
	}
	if params.URL == "" {
		respondWithError(w, 403, fmt.Sprintf("No URL included: %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      user.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}
