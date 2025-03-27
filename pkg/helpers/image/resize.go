package image

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

type ResizeImageType int

const (
	Thumb ResizeImageType = iota
	Mid
	Orig
)

type ResizeConfig struct {
	FullPath     string
	ImgWidth     uint
	ImgType      ResizeImageType
	CompressAlgo resize.InterpolationFunction
	ForceToJpeg  bool
}

func ResizeImage(params ResizeConfig) error {
	file, err := os.Open(params.FullPath)
	if err != nil {
		return err
	}

	fmt.Printf("Filename: %s\n", file.Name())
	fileExt := strings.Split(file.Name(), ".")[1]
	var img image.Image

	if fileExt == "png" {
		img, err = png.Decode(file)
	} else {
		img, err = jpeg.Decode(file)
	}

	if err != nil {
		return err
	}
	file.Close()

	m := resize.Resize(params.ImgWidth, 0, img, params.CompressAlgo)

	var fileCategoryName string
	if params.ImgType == Thumb {
		fileCategoryName = "thumb"
	} else {
		fileCategoryName = "midsize"
	}

	finalFileExt := fileExt
	if params.ForceToJpeg && !(fileExt == "jpg" || fileExt == "jpeg") {
		finalFileExt = "jpeg"
	}

	fullPathEls := strings.Split(params.FullPath, "/")
	fullPathDir := strings.Join(fullPathEls[:len(fullPathEls)-1], "/")
	outFullPath := fmt.Sprintf("%s/%s.%s", fullPathDir, fileCategoryName, finalFileExt)

	fmt.Printf("outFullPath: %s\n", outFullPath)
	out, err := os.Create(outFullPath)
	if err != nil {
		return err
	}
	defer out.Close()

	if fileExt == "png" && !params.ForceToJpeg {
		png.Encode(out, m)
	} else {
		jpeg.Encode(out, m, nil)
	}

	return nil
}
