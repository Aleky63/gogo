package worker

import (
	"errors"
	"yourgo/networkerr"
)

var ErrServiceUnavailable = networkerr.New("service unavailable", 222)
var ErrNoInternet error = errors.New("no internet")

type Worker struct{}

func New() *Worker {
	return &Worker{}
}
