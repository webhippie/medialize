package util

import (
	"strings"
)

// TrimPrefix can drop a specific prefix.
func TrimPrefix(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		s = s[len(prefix):]
	}

	return s
}

// TrimSuffix can drop a specific suffix.
func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}

	return s
}
