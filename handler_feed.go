package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BaneleJerry/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) createFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	var newFeed struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&newFeed)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      newFeed.Name,
		Url:       newFeed.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create feed")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}

func (cfg *apiConfig) getFeedsHandler(w http.ResponseWriter, r *http.Request) {
    feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not get Feeds")
	}
	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}