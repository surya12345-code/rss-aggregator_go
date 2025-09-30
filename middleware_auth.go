package main

import (
	"fmt"
	"net/http"

	"github.com/surya123/RSSaggregator/internal/auth"
	"github.com/surya123/RSSaggregator/internal/database"
)

type authhandler func(http.ResponseWriter, *http.Request, database.User)

func (apicfg *apiconfig) middlewareAuth(handler authhandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondwitherror(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apicfg.DB.Getuserbyapikey(r.Context(), apikey)

		if err != nil {
			respondwitherror(w, 400, fmt.Sprintf("couldn't get user: %v", err))
			return

		}
		handler(w, r, user)
	}

}
