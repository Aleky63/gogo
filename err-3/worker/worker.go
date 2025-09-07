package worker

import (
	"err-3/networkerr"
	"errors"
)

var ErrServiceUnavailable = networkerr.New("service unavailable", 2225566)
var ErrNoInternet error = errors.New("no internet")

type Worker struct{}

func New() *Worker {
	return &Worker{}
}
func (Worker) DoWork() error {
	// return ErrServiceUnavailable
	return networkerr.NewWithErr(ErrNoInternet, 235)
	// return errors.New("file not found")
}

// https://stepik.org/lesson/1501026/step/1?unit=1521142
// Оборачивание пользовательских ошибок
