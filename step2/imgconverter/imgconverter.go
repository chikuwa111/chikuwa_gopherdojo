// Package imgconverter provides Decode and Encode functions.
// These functions need file path.
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

// Decode does decode image in specific path.
// This supports jpg(jpeg) and png.
func Decode(path string) (Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return Image{nil}, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return Image{nil}, err
	}

	return Image{img}, nil
}

// Encode does encode image into specific format and create a file.
// This supports jpg(jpeg) and png.
func (img *Image) Encode(dest string) error {
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
