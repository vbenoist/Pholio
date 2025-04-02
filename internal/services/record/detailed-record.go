package record

import (
	"github.com/jackidu14/pholio/internal/helpers/file"
	apimodels "github.com/jackidu14/pholio/internal/models/api"
	imagehelper "github.com/jackidu14/pholio/pkg/helpers/image"
)

func GetDetailedRecords(pgParams apimodels.PaginationQuery) (apimodels.PaginatedResults[apimodels.DetailedRecord], error) {
	var results apimodels.PaginatedResults[apimodels.DetailedRecord]

	paginatedResult, err := GetRecords(pgParams)
	if err != nil {
		return results, err
	}

	for _, doc := range paginatedResult.Documents {
		path, err := file.GetFileFullpath(doc.Id.Hex(), imagehelper.Orig)
		if err != nil {
			return results, err
		}

		results.Documents = append(results.Documents, apimodels.DetailedRecord{
			Record: doc,
			Folder: path,
		})
	}

	results.Pagination = paginatedResult.Pagination
	return results, nil
}
