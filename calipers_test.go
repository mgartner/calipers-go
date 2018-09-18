package calipers

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	path                string
	expectedMeasurement Measurement
}

var digitsRegexp = regexp.MustCompile(`\d+`)

func generateTests(path string, t ImageType) ([]testCase, error) {
	var tc []testCase

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return tc, err
	}

	for _, f := range files {
		splits := digitsRegexp.FindAllString(f.Name(), 2)

		width, err := strconv.Atoi(splits[0])
		if err != nil {
			return tc, err
		}

		height, err := strconv.Atoi(splits[1])
		if err != nil {
			return tc, err
		}

		m := Measurement{t, width, height}
		filePath := filepath.Join(path, f.Name())

		tc = append(tc, testCase{filePath, m})
	}

	return tc, nil
}

func TestValidGIF(t *testing.T) {
	tests, err := generateTests("assets/gif", GIF)
	assert.NoError(t, err)

	for _, tt := range tests {
		result, err := Measure(tt.path)

		assert.NoError(t, err)
		assert.Equal(t, result, tt.expectedMeasurement)
	}
}

func TestValidPNG(t *testing.T) {
	tests, err := generateTests("assets/png", PNG)
	assert.NoError(t, err)

	for _, tt := range tests {
		result, err := Measure(tt.path)

		assert.NoError(t, err)
		assert.Equal(t, result, tt.expectedMeasurement)
	}
}

func TestUnknownFile(t *testing.T) {
	_, err := Measure("assets/jpeg/114x118.jpg")

	assert.EqualError(t, err, "unknown file type")
}
