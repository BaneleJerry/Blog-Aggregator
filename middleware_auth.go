package main

import (
	"net/http"

	"github.com/BaneleJerry/Blog-Aggregator/internal/auth"
	"github.com/BaneleJerry/Blog-Aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(&r.Header)

		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		user, err := cfg.DB.GetUSerBYApi(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Could not get user")
			return
		}
		handler(w, r, user)
	}
}
