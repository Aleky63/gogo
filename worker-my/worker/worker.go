package worker

import (
	"errors"
)

var ErrorFileNotExists = errors.New("file not exists")

type Worker struct {
}

func New() *Worker {
	return &Worker{}
}

func (Worker) DoWork(path string) error {
	return NewFileNotExistsError(path)

}
