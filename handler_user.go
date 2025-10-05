package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/surya123/RSSaggregator/internal/database"
)

func (apicfg *apiconfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondwitherror(w, 400, fmt.Sprintf("Error Parsing json :%s", err))
		return
	}
	user, err := apicfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondwitherror(w, 400, fmt.Sprintf("coldn't respond user :%v", err))
	}

	respondWithJson(w, 201, databaseUserToUser(user))

}
func (apicfg *apiconfig) handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJson(w, 200, databaseUserToUser(user))
}
func (apicfg *apiconfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apicfg.DB.GetPostsForUsers(r.Context(), database.GetPostsForUsersParams{
		UserID: user.ID,
		Limit:  50,
	})
	if err != nil {
		respondwitherror(w, 400, fmt.Sprintf("Error getting posts :%v", err))
		return
	}

	respondWithJson(w, 200, databasepoststoposts(posts))
}
