package image

import (
	"log"

	"github.com/nfnt/resize"
)

func ResizeImageMid(fullPath string) {
	resizeParams := ResizeConfig{
		FullPath:     fullPath,
		ImgWidth:     800,
		ImgType:      Mid,
		CompressAlgo: resize.Lanczos3,
		ForceToJpeg:  false,
	}

	err := resizeImage(resizeParams)

	if err != nil {
		log.Printf("helpers:image::ResizeImageMid an error occured: %s", err.Error())
	}
}
