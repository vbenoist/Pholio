package apimodels

import databasemodels "github.com/jackidu14/pholio/internal/models/database"

type GroupbyRecord struct {
	GroupBy string                  `json:"group-by"`
	Results []databasemodels.Record `json:"results"`
}
