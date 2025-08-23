package project

import (
	"errors"
	"github.com/google/uuid"
)

type Project struct {
	ID    uuid.UUID
	Name  string
	Tasks []Task
}

func New(ID uuid.UUID, name string) (*Project, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Project{
		ID:    ID,
		Name:  name,
		Tasks: make([]Task, 0),
	}, nil
}

func AddTask(p *Project)(task Task)  error {

}
