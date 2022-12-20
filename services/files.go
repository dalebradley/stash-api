package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dalebradley/stash-api/models"
	"github.com/dalebradley/stash-api/storage"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-playground/validator/v10"

	"honnef.co/go/tools/config"
)

// FilesService contains the Repository for db access
type FilesService struct {
	Repository storage.Mongo
	Config     config.Config
}

func CreateFile(req *http.Request, createFileRequest models.FileResourceRest) (*models.FileResourceRest, int, error) {
	// Add trace logs
	// TODO: Hardcoded createdBy until auth implemented
	// createdBy := models.CreatedBy{
	// 	Email:    "deebrad@hotmail.co.uk",
	// 	Forename: "Dale",
	// 	Surname:  "Bradley",
	// 	ID:       "1",
	// }
	// fileResourceDB := models.FileResourceDB{
	// 	Type:      "test",
	// 	ID:        "1",
	// 	CreatedAt: time.Now(),
	// 	CreatedBy: createdBy,
	// }

	err := validateCreateFileRequestBody(createFileRequest)
	spew.Dump("AFTER VALIDATE")
	spew.Dump(err)
	if err != nil {
		spew.Dump("In ERROR")
		err = fmt.Errorf("invalid file resource: [%v]", err)
		//LOG ERROR
		return nil, 400, err
	}

	err = storage.CreateFileResource(createFileRequest)
	if err != nil {
		err = fmt.Errorf("error creating file resource: [%v]", err)
		//LOG ERROR
		return nil, 500, err
	}
	fileResourseRest := models.FileResourceRest(createFileRequest)
	return &fileResourseRest, 200, nil
}

func GetFile(req *http.Request, id string) (models.FileResourceRest, int, error) {
	fileResourceDB, err := storage.GetFileResource(id)
	if err != nil {
		//Handle
	}
	return models.FileResourceRest(*fileResourceDB), 200, nil
}

func validateCreateFileRequestBody(createFileRequest models.FileResourceRest) error {
	validate := validator.New()
	spew.Dump("In VALIDATE")
	spew.Dump(createFileRequest)
	err := validate.Struct(createFileRequest)
	if err != nil {
		return err
	}
	// TODO: Check all fields
	if createFileRequest.ID == "" {
		return errors.New("empty ID")
	}
	if createFileRequest.Amount == 0 {
		return errors.New("no amount provided")
	}
	if createFileRequest.Type == "" {
		return errors.New("no type provided")
	}
	return nil
}
