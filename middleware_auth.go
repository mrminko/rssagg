package main

import (
	"fmt"
	"github.com/mrminko/rssagg/internal/auth"
	"github.com/mrminko/rssagg/internal/database"
	"net/http"
)

type authedHandler func(w http.ResponseWriter, response *http.Request, user database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Authentication failure: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Error when extracting user: %v", err))
			return
		}
		handler(w, r, user)
	}

}
