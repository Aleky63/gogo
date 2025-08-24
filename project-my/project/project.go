package project

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Project struct {
	ID      uuid.UUID
	Name    string
	TaskIds []uuid.UUID
	Tasks   map[uuid.UUID]Task
}

func New(ID uuid.UUID, name string) (*Project, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Project{
		ID:      ID,
		Name:    name,
		TaskIds: make([]uuid.UUID, 0),
		Tasks:   make(map[uuid.UUID]Task),
	}, nil
}

func (p *Project) AddTask(task Task) error {

	if _, ok := p.Tasks[task.ID]; ok {
		return errors.New(" task already exists")

	}
	p.TaskIds = append(p.TaskIds, task.ID)
	p.Tasks[task.ID] = task
	return nil
}

func (p *Project) UpdateTask(task Task) error {
	if _, ok := p.Tasks[task.ID]; !ok {
		return errors.New(" task not exists")
	}
	p.Tasks[task.ID] = task
	return nil
}

func (p Project) FilterTasksByStatus(status Status) []Task {
	var res []Task
	for _, tid := range p.TaskIds {
		task := p.Tasks[tid]
		if task.Status == status {
			res = append(res, task)
		}
	}
	return res
}

func (p Project) PrintInfo() {
	fmt.Printf("Проект: %s\n", p.Name)
	for _, tid := range p.TaskIds {
		task := p.Tasks[tid]
		fmt.Printf(" Задача: %s (ID: %s), Описание: %s, Статус: %s\n", task.Title, task.ID, task.Describe, task.Status)
	}
	fmt.Println("---")
}
