// Package calipers measures the dimensions of image files
// quickly by not loading the entire image into memory.
package calipers

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"os"
)

// Measurement struct contains information about the type and
// dimensions (in pixels) of an image.
type Measurement struct {
	Type   ImageType
	Width  int
	Height int
}

// ImageType denotes a type of image, such as PNG.
type ImageType string

const (
	// PNG format.
	PNG ImageType = "png"
	// JPEG format.
	JPEG ImageType = "jpeg"
)

var pngHeader = []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}

// Measure returns the dimensions of an image file.
func Measure(path string) (Measurement, error) {

	file, err := os.Open(path)
	if err != nil {
		return Measurement{}, err
	}

	imageType, err := detect(file)
	if err != nil {
		return Measurement{}, err
	}

	switch imageType {
	case PNG:
		return measurePNG(file)
	default:
		return Measurement{}, errors.New("unable to measure file")
	}
}

func detect(file *os.File) (ImageType, error) {
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}

	buffer := make([]byte, 8)
	if _, err := io.ReadFull(file, buffer); err != nil {
		return "", err
	}

	switch {
	case bytes.Equal(buffer, pngHeader):
		return PNG, nil
	default:
		return "", errors.New("unknown file type")
	}
}

func measurePNG(file *os.File) (Measurement, error) {
	_, err := file.Seek(16, 0)
	if err != nil {
		return Measurement{}, err
	}

	buffer := make([]byte, 8)
	if _, err := io.ReadFull(file, buffer); err != nil {
		return Measurement{}, errors.New("unable to read PNG")
	}

	width := binary.BigEndian.Uint32(buffer[0:4])
	height := binary.BigEndian.Uint32(buffer[4:8])

	return Measurement{PNG, int(width), int(height)}, nil
}
