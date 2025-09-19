package main

import "net/http"

func handlererror(w http.ResponseWriter, r *http.Request) {
	respondwitherror(w, 400, "There is something error")

}
