package main

import (
	"fmt"
	"net/http"
)

func (apicfg *apiconfig) Getallusers(w http.ResponseWriter, r *http.Request) {
	users, err := apicfg.DB.Getusers(r.Context())
	if err != nil {
		respondwitherror(w, 400, fmt.Sprintf("Error getting users: %v", err))
		return
	}

	respondWithJson(w, 201, databaseGetuserstousers(users))

}
