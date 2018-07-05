package photo

import (
	// "fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	// "github.com/metakeule/fmtdate"
	"github.com/webhippie/medialize/pkg/util"
)

var (
	// validExtensions are used to validate a file extension.
	validExtensions = []string{
		".png",
		".gif",
		".jpg",
		".jpeg",
	}
)

// File represents a single found source file.
type File struct {
	Path               string
	CalculatedChecksum string
	CalculatedCreation time.Time
	Info               os.FileInfo
}

// String converts a file into a string representation.
func (h *File) String() string {
	return h.Path
}

// Valid checks if the file got a valid extension.
func (h *File) Valid() bool {
	ext := strings.ToLower(filepath.Ext(h.Info.Name()))

	for _, check := range validExtensions {
		if check == ext {
			return true
		}
	}

	return false
}

// Ext tries to return a cleaned file extension.
func (h *File) Ext() string {
	result := strings.ToLower(filepath.Ext(h.Info.Name()))

	switch result {
	case ".jpeg":
		return ".jpg"
	default:
		return result
	}
}

// Checksum generates a SHA256 checksum from file content.
func (h *File) Checksum() (string, error) {
	result, err := util.Checksum(h.Path)

	if err != nil {
		return "", err
	}

	h.CalculatedChecksum = result
	return result, nil
}
