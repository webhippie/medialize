package util

import (
	"path/filepath"
	"strings"
)

func CleanExt(file string) string {
	result := strings.ToLower(
		filepath.Ext(file))

	switch result {
	case ".jpeg":
		return ".jpg"
	default:
		return result
	}
}
