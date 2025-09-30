package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/surya123/RSSaggregator/internal/database"
)

func (apicfg *apiconfig) handlerFeedfollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondwitherror(w, 400, fmt.Sprintf("Error Parsing json :%s", err))
		return
	}
	feedfollow, err := apicfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		respondwitherror(w, 400, fmt.Sprintf("coldn't create feed  :%v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedfollowstofeedfollows(feedfollow))

}
func (apicfg *apiconfig) handlerGetFeedfollows(w http.ResponseWriter, r *http.Request, user database.User) {
	userfeeds, err := apicfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondwitherror(w, 400, fmt.Sprintf("coldn't get feeds  :%v", err))
		return
	}

	respondWithJson(w, 200, databaseGetuserfeedstofeeds(userfeeds))

}
