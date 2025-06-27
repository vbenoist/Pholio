package apimodels

import databasemodels "github.com/vbenoist/pholio/internal/models/database"

type DetailedRecord struct {
	databasemodels.Record
	Folder string `json:"folder"`
	Pin    bool   `json:"pin"`
}
