package record

import (
	"github.com/vbenoist/pholio/internal/helpers/file"
	apimodels "github.com/vbenoist/pholio/internal/models/api"
	databasemodels "github.com/vbenoist/pholio/internal/models/database"
	imagehelper "github.com/vbenoist/pholio/pkg/helpers/image"
)

func GetDetailedRecords(pgParams apimodels.PaginationQuery) (apimodels.PaginatedResults[apimodels.DetailedRecord], error) {
	var results apimodels.PaginatedResults[apimodels.DetailedRecord]

	paginatedResult, err := GetRecordsStars(pgParams)
	if err != nil {
		return results, err
	}

	for _, doc := range paginatedResult.Documents {
		path, err := file.GetFileFullpath(doc.Record.Id.Hex(), imagehelper.Orig)
		if err != nil {
			return results, err
		}

		results.Documents = append(results.Documents, apimodels.DetailedRecord{
			Record: databasemodels.Record{
				Id:          doc.Id,
				Description: doc.Description,
				Location:    doc.Location,
				Date:        doc.Date,
			},
			Folder: path,
			Pin:    len(doc.Pin) > 0,
		})
	}

	results.Pagination = paginatedResult.Pagination
	return results, nil
}
