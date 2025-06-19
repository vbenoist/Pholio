package apimodels

import databasemodels "github.com/jackidu14/pholio/internal/models/database"

type DetailedRecord struct {
	databasemodels.Record
	Folder string `json:"folder"`
	Pin    bool   `json:"pin"`
}
