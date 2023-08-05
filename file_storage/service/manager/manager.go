package manager

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileManager struct {
	directory string
}

func NewFileManager(storageDirectory string) (*FileManager, error) {
	dir, err := filepath.Abs(storageDirectory)
	if err != nil {
		return &FileManager{}, fmt.Errorf("invalid directory path provided: %s", storageDirectory)
	}

	manager := &FileManager{
		directory: dir,
	}

	// validate directory exists with write access
	_, err = os.Stat(manager.directory)
	if os.IsNotExist(err) {
		// attempt to create
		err = os.MkdirAll(manager.directory, 0755)
		if err != nil {
			return manager, err
		}
	} else {
		return manager, err
	}

	// check if writable
	file, err := os.CreateTemp(manager.directory, "tmp")
	if err != nil {
		return manager, fmt.Errorf("provided directory does not have the required permissions. directory must be readable & writable")
	}

	defer os.Remove(file.Name())
	defer file.Close()

	return manager, nil
}

func (fm *FileManager) ListAllFiles(path string) ([]File, error) {
	fullPath := filepath.Join(fm.directory, path)
	f, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}

	filesInfo, err := f.ReadDir(0)
	if err != nil {
		return nil, err
	}

	files := make([]File, 0)

	for _, v := range filesInfo {
		info, err := v.Info()
		if err != nil {
			continue
		}

		files = append(files, File{
			Name:         v.Name(),
			Size:         info.Size(),
			IsDir:        v.IsDir(),
			LastModified: info.ModTime(),
		})
	}

	return files, nil
}
