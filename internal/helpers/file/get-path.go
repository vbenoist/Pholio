package file

import (
	"fmt"
	"os"
	"strings"

	"github.com/jackidu14/pholio/internal/helpers/cfg"
	"github.com/jackidu14/pholio/internal/helpers/image"
	"github.com/pkg/errors"
)

func GetFileFullpath(recordId string, fileType image.ResizeImageType) (string, error) {
	config := cfg.SetServerConfig()
	folderPath := fmt.Sprintf("%s/%s/", config.FileManager.UploadPath, recordId)

	folder, err := os.ReadDir(folderPath)
	if err != nil {
		return "", err
	}

	var fileName string
	for i := 0; i < len(folder); i++ {
		folderFileCatg := strings.Split(folder[i].Name(), ".")[0]

		if fileType == image.Thumb && folderFileCatg == "thumb" {
			fileName = folder[i].Name()
		} else if fileType == image.Mid && folderFileCatg == "midsize" {
			fileName = folder[i].Name()
		} else if fileType == image.Orig && folderFileCatg == "original" {
			fileName = folder[i].Name()
		}
	}

	if fileName == "" {
		return "", errors.New("No corresponding file has been found")
	}

	return folderPath + fileName, nil
}
