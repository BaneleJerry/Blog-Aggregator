package main

import (
	"net/http"
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
