package photo

import (
	"fmt"
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

// CreationTime tries to fetch the creation date from EXIF.
func CreationTime(file string) (time.Time, error) {
	handle, err := os.Open(file)

	if err != nil {
		return time.Time{}, fmt.Errorf("Failed to open file: %s", err)
	}

	info, err := exif.Decode(handle)

	if err != nil {
		return time.Time{}, fmt.Errorf("Failed to parse file: %s", err)
	}

	taken, err := info.DateTime()

	if err != nil {
		return time.Time{}, fmt.Errorf("Failed to get time: %s", err)
	}

	return taken, nil
}
