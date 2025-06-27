package apimodels

import databasemodels "github.com/vbenoist/pholio/internal/models/database"

type GroupbyRecord struct {
	GroupBy string                  `json:"group-by"`
	Results []databasemodels.Record `json:"results"`
}
