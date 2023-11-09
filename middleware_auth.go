package main

import (
	"RossDooney/blog-aggregator/internal/auth"
	"RossDooney/blog-aggregator/internal/database"
	"fmt"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apikey)

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user", err))
			return
		}

		handler(w, r, user)

	}
}
