// +build linux darwin

package photo

import (
	// "fmt"
	"time"
	// "github.com/xiam/exif"
)

// Creation tries to extract the correct creation time.
func (h *File) Creation() (time.Time, error) {
	// data, err := exif.Read(h.Path)

	// if err != nil {
	// 	return time.Now(), fmt.Errorf("Failed to read file. %s", err)
	// }

	// taken, ok := data.Tags["Date and Time"]

	// if !ok {
	// 	return time.Now(), fmt.Errorf("Time attribute doesn't exist")
	// }

	// parsed, err := time.Parse("2006:01:02 15:04:05", taken)

	// if err != nil {
	// 	return time.Now(), fmt.Errorf("Failed to parse time. %s", err)
	// }

	// h.CalculatedCreation = parsed
	// return parsed, nil

	return time.Now(), nil
}
