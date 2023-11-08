package main

import (
	"RossDooney/blog-aggregator/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	parms := parameters{}

	err := decoder.Decode(&parms)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json", err))
		return
	}

	apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      parms.Name,
	})

	respondWithJSON(w, 200, struct{}{})
}
