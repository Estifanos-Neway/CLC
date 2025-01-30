package dir

import (
	"os"
)

func GetDirectoryContents(dir string) (*[]string, error) {
	d, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	dirContents := []string{}
	for _, dc := range d {
		dirContents = append(dirContents, dc.Name())
	}
	return &dirContents, nil
}
