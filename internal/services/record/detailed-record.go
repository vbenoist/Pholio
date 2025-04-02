package record

import (
	"github.com/jackidu14/pholio/internal/helpers/file"
	apimodels "github.com/jackidu14/pholio/internal/models/api"
	imagehelper "github.com/jackidu14/pholio/pkg/helpers/image"
)

func GetDetailedRecords() ([]apimodels.DetailedRecord, error) {
	var results []apimodels.DetailedRecord
	documents, err := GetRecords()
	if err != nil {
		return results, err
	}

	for i := 0; i < len(documents); i++ {
		path, err := file.GetFileFullpath(documents[i].Id.Hex(), imagehelper.Orig)
		if err != nil {
			return results, err
		}

		results = append(results, apimodels.DetailedRecord{
			Record: documents[i],
			Folder: path,
		})
	}

	return results, nil
}
