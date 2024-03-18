package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/BaneleJerry/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeedFollower(w http.ResponseWriter, r *http.Request, user database.User) {
	var param struct {
		Feed_id uuid.UUID `json:"feed_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	feed_follow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    param.Feed_id,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create feed_follow")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feed_follow))
}

func (cfg *apiConfig) handlerDeleteFeedFollower(w http.ResponseWriter, r *http.Request, user database.User) {
	requestURL := r.URL.String()
	reuestSplit := strings.Split(requestURL, "/")

	if len(reuestSplit) < 4 {
		respondWithError(w, http.StatusBadRequest, "Invalid URL")
		return
	}

	feedFollowStr := reuestSplit[3]
	if feedFollowStr == "" {
		respondWithError(w, http.StatusBadRequest, "feedFollowID Empty")
		return
	}
	feedFollowID, err := uuid.Parse(feedFollowStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feedFollowID")
		return
	}
	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not delete feed_follow")
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (cfg *apiConfig) handlerGetFeedFollowers(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollows, err := cfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not get Feeds")
	}
	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedsFollows))
}
