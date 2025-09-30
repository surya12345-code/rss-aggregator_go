package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"github.com/surya123/RSSaggregator/internal/database"
)

func (apicfg *apiconfig) handlerCreatefeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondwitherror(w, 400, fmt.Sprintf("Error Parsing json :%s", err))
		return
	}
	feed, err := apicfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.Url,
		Name:      params.Name,
		UserID:    user.ID,
	})
	if err != nil {
		respondwitherror(w, 400, fmt.Sprintf("coldn't create feed  :%v", err))
		return
	}

	respondWithJson(w, 201, databasefeedtofeed(feed))

}
func (apicfg *apiconfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apicfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondwitherror(w, 400, fmt.Sprintf("Error getting feeds :%v", err))
		return
	}

	respondWithJson(w, 201, databasefeedstofeeds(feeds))

}
func (apicfg *apiconfig) handlerDeletefeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedfollowdelete := chi.URLParam(r, "feedfollowid")
	feedfollowid, err := uuid.Parse(feedfollowdelete)
	if err != nil {
		respondwitherror(w, 400, fmt.Sprintf("Couldn't parse feedfollowid:%v", err))
		return
	}
	err = apicfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedfollowid,
		UserID: user.ID,
	})
	if err != nil {
		respondwitherror(w, 400, fmt.Sprintf("Couldn't delete feedfollowid:%v", err))
		return
	}
	respondWithJson(w, 201, struct{}{})

}
