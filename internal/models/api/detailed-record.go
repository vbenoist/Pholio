package apimodels

import databasemodels "github.com/jackidu14/pholio/internal/models/database"

type DetailedRecord struct {
	databasemodels.Record
	Folder string
}
