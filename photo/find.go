package photo

import (
	"os"
	"path/filepath"
)

// FindFiles walkes through the file tree and fetches files.
func FindFiles(searchPath string) ([]string, error) {
	fileList := []string{}

	err := filepath.Walk(
		searchPath,
		func(path string, f os.FileInfo, err error) error {
			if f.IsDir() {
				return nil
			}

			fileList = append(
				fileList,
				path)

			return nil
		})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}
