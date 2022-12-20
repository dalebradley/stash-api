package routes

import (
	"github.com/dalebradley/stash-api/handlers"

	"github.com/gorilla/mux"
)

func Register(mainRouter *mux.Router) {
	r := mainRouter.PathPrefix("/stash").Subrouter()
	r.HandleFunc("/", handlers.HandleGetFiles).Methods("GET")
	r.HandleFunc("/create", handlers.HandleCreateFile).Methods("POST")
	r.HandleFunc("/healthcheck", handlers.HandleHealthcheck).Methods("GET")
	r.HandleFunc("/{id}", handlers.HandleGetFile).Methods("GET")
}
