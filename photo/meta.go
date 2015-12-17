package photo

import (
	"errors"
	"github.com/rwcarlsen/goexif/exif"
	"os"
	"time"
)

func CreationTime(file string) (time.Time, error) {
	handle, err := os.Open(file)

	if err != nil {
		return time.Time{}, errors.New("Failed to open file")
	}

	info, err := exif.Decode(handle)

	if err != nil {
		return time.Time{}, errors.New("Failed to parse file")
	}

	taken, err := info.DateTime()

	if err != nil {
		return time.Time{}, errors.New("Failed to get time")
	}

	return taken, nil
}
