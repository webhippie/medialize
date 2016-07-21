package photo

import (
	"path/filepath"
	"strings"
)

// ValidExtension detects valid photo file extensions.
func ValidExtension(file string) bool {
	valid := []string{
		".png",
		".gif",
		".jpg",
		".jpeg",
	}

	ext := strings.ToLower(
		filepath.Ext(file))

	for _, check := range valid {
		if check == ext {
			return true
		}
	}

	return false
}
