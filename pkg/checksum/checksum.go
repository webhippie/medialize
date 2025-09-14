package checksum

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// New generates a SHA256 checksum from file content.
func New(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	defer func() { _ = file.Close() }()

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("failed to copy file: %w", err)
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
