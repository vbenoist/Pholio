package image

import (
	"log"

	"github.com/nfnt/resize"
)

func ResizeImageThumb(fullPath string) {
	resizeParams := ResizeConfig{
		FullPath:     fullPath,
		ImgWidth:     400,
		ImgType:      Thumb,
		CompressAlgo: resize.Lanczos3,
		ForceToJpeg:  true,
	}

	err := resizeImage(resizeParams)

	if err != nil {
		log.Printf("helpers:image::ResizeImageThumb an error occured: %s", err.Error())
	}
}
