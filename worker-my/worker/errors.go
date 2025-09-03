package worker

import "fmt"

type FileNotExistsError struct {
	Path string
}

func NewFileNotExistsError(path string) *FileNotExistsError {
	return &FileNotExistsError{
		Path: path,
	}
}

func (e FileNotExistsError) Error() string {
	return fmt.Sprintf("path %s is not exists", e.Path)
}
