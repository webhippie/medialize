//go:build !linux && !darwin
// +build !linux,!darwin

package photo

import (
	"fmt"
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

// Creation tries to extract the correct creation time.
func (h *File) Creation() (time.Time, error) {
	handle, err := os.Open(h.Path)

	if err != nil {
		return time.Now(), fmt.Errorf("Failed to open file. %s", err)
	}

	info, err := exif.Decode(handle)

	if err != nil {
		return time.Now(), fmt.Errorf("Failed to parse exif. %s", err)
	}

	parsed, err := info.DateTime()

	if err != nil {
		return time.Now(), fmt.Errorf("Failed to parse time. %s", err)
	}

	h.CalculatedCreation = parsed
	return parsed, nil
}
