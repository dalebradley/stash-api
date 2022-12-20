package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dalebradley/stash-api/services"

	"github.com/dalebradley/stash-api/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func HandleGetFiles(w http.ResponseWriter, r *http.Request) {
	// TODO pagination?
	w.WriteHeader(http.StatusOK)
}

func HandleCreateFile(w http.ResponseWriter, r *http.Request) {
	requestDecoder := json.NewDecoder(r.Body)
	var incomingCreateFileRequest models.IncomingCreateFileRequest
	err := requestDecoder.Decode(&incomingCreateFileRequest)
	if err != nil {
		err = fmt.Errorf("error decoding incoming request when creating file: %s", err)
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	createdBy := models.CreatedBy{
		Email:    "deebrad@hotmail.co.uk",
		Forename: "Dale",
		Surname:  "Bradley",
		ID:       "1",
	}

	fileResourseRest := models.FileResourceRest{
		ID:        uuid.New().String(),
		Type:      incomingCreateFileRequest.Type,
		CreatedAt: time.Now(),
		CreatedBy: createdBy,
	}
	_, status, err := services.CreateFile(r, fileResourseRest)
	if status == 400 {
		err = fmt.Errorf("error creating file: %s", err)
		fmt.Println(err)
		w.WriteHeader((http.StatusBadRequest))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func HandleGetFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileResoureRest, status, err := services.GetFile(r, vars["id"])
	if status == 400 {
		err = fmt.Errorf("error getting file <%s> with err: %s", vars["id"], err)
		fmt.Println(err)
		w.WriteHeader((http.StatusBadRequest))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(fileResoureRest)
	if err != nil {
		err = fmt.Errorf("error writing response for GET file: %s", err)
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
