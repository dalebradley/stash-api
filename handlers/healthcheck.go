package handlers

import (
	"net/http"
)

func HandleHealthcheck(w http.ResponseWriter, r *http.Request) {
	// TODO
	// Can connect to mongo?
	w.WriteHeader(http.StatusOK)
}
