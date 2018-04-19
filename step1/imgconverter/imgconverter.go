// Package imgconverter provides Convert function.
package imgconverter

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// Image is a wrapper of image.Image
type Image struct {
	image.Image
}

// Convert does convert image into specific format and create a file.
// This supports jpg(jpeg) and png.
func (img *Image) Convert(dest string) error {
	file, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer file.Close()

	switch filepath.Ext(dest) {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
	case ".png":
		err = png.Encode(file, img)
	default:
		err = errors.New("invalid dest")
	}
	return err
}
