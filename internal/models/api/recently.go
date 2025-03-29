package apimodels

import (
	databasemodels "github.com/jackidu14/pholio/internal/models/database"
)

// Difference between Lastly / Lately is just from configured date
type RecentlyRecords struct {
	Lastly []databasemodels.Record `json:"lastly"`
	Lately []databasemodels.Record `json:"lately"`
}
