package models

import "time"

// FileResourceDB contains details of a file
type FileResourceDB struct {
	ID        string    `bson:"id"`
	Type      string    `bson:"type"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
	CreatedBy CreatedBy `bson:"created_by" json:"created_by"`
}

// CreatedByDB is the user who is creating the file
type CreatedBy struct {
	Email    string `bson:"email" json:"email"`
	Forename string `bson:"forename" json:"forename"`
	ID       string `bson:"id" json:"id"`
	Surname  string `bson:"surname" json:"surname"`
}

// File contains details of a file
type FileResourceRest struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CreatedBy CreatedBy `json:"created_by"`
}

// IncomingCreateFileRequest represents the incoming request when creating a file
type IncomingCreateFileRequest struct {
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
}

// Files is a list of File items
type FilesRest struct {
	Files []FileResourceRest `json:"files"`
}
