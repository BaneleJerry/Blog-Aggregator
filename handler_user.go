package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BaneleJerry/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {

	var newUser struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      newUser.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create user")
		return
	}
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

func (cfg *apiConfig) getUserHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
