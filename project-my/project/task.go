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
