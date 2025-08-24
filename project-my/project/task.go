package project

import (
	"errors"

	"github.com/google/uuid"
)

type Status string

const (
	StatusActive = "ACTIVE"
	StatusClosed = "CLOSED"
)

type Task struct {
	ID       uuid.UUID
	Title    string
	Describe string
	Status   Status
}

func NewTask(ID uuid.UUID, title, describe string) (*Task, error) {
	if title == "" {
		return nil, errors.New("empty title")
	}
	if describe == "" {
		return nil, errors.New("empty describe")
	}
	return &Task{
		ID:       ID,
		Title:    title,
		Describe: describe,
		Status:   StatusActive,
	}, nil
}
func (t *Task) UpdateDescription(desc string) error {
	if desc == "" {
		return errors.New("empty description")

	}
	if t.Status != StatusActive {
		return errors.New("task status not active")
	}
	t.Describe = desc
	return nil
}
func (t *Task) Close() error {

	if t.Status == StatusClosed {
		return errors.New("status closed already")
	}
	t.Status = StatusClosed
	return nil
}
