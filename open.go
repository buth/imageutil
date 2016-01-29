package imageutil

import (
	"image"
	"os"
)

// OpenImage opens the named image file, decodes it, and closes it, returning
// an image and any error. BYO decoder support.
func OpenImage(name string) (image.Image, string, error) {
	imageFile, err := os.Open(name)
	if err != nil {
		return nil, "", err
	}

	imageFileDecoded, format, err := image.Decode(imageFile)
	if err != nil {
		return nil, "", err
	}

	if err := imageFile.Close(); err != nil {
		return nil, "", err
	}

	return imageFileDecoded, format, nil
}
