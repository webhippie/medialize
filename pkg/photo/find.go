package photo

import (
	"os"
	"path/filepath"
)

// Find walkes through the file tree and fetches files.
func Find(search string) ([]*File, error) {
	list := make([]*File, 0)

	err := filepath.Walk(search, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}

		list = append(
			list,
			&File{
				Path: path,
				Info: f,
			},
		)

		return err
	})

	if err != nil {
		return nil, err
	}

	return list, nil
}
