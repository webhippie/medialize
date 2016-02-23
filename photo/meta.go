package photo

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

func CreationTime(file string) (time.Time, error) {
	handle, err := os.Open(file)

	if err != nil {
		return time.Time{}, errors.New(
			fmt.Sprintf("Failed to open file: %s", err))
	}

	info, err := exif.Decode(handle)

	if err != nil {
		return time.Time{}, errors.New(
			fmt.Sprintf("Failed to parse file: %s", err))
	}

	taken, err := info.DateTime()

	if err != nil {
		return time.Time{}, errors.New(
			fmt.Sprintf("Failed to get time: %s", err))
	}

	return taken, nil
}
