package main

import (
	"log"
	"net/http"

	"github.com/dalebradley/stash-api/conf"
	"github.com/dalebradley/stash-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	cfg := conf.Get()

	mainRouter := mux.NewRouter()
	mainRouter.StrictSlash(true)
	routes.Register(mainRouter)

	log.Println("Starting server on port " + cfg.StashPort[1:])
	err := http.ListenAndServe(cfg.StashPort, mainRouter)
	if err != nil {
		log.Fatal("Error starting server", err)
	}

}
