package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BaneleJerry/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

//Helpers

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	type readiness struct {
		Status string `json:"status"`
	}
	status := readiness{Status: "Ok"}
	respondWithJSON(w, http.StatusOK, status)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Internal Server Error"
	respondWithError(w, http.StatusInternalServerError, msg)
}

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {

	var newUser struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      newUser.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError,err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}