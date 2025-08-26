package main

import (
	"errors"
	"fmt"
)

type Character interface {
	Attack() string
}
type Warrior struct {
	name string
}

func NewWarrior(name string) (*Warrior, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Warrior{name: name}, nil
}
func (w *Warrior) Attack() string {
	return fmt.Sprintf("Воин %s бьет мечом.", w.name)
}

type Mage struct {
	name string
}

func NewMage(name string) (*Mage, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Mage{name: name}, nil
}

func (m *Mage) Attack() string {
	return fmt.Sprintf("Маг %s колдует огненный шар.", m.name)
}

type Archer struct {
	name string
}

func NewArcher(name string) (*Archer, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}
	return &Archer{name: name}, nil
}

func (a *Archer) Attack() string {
	return fmt.Sprintf("Лучник %s выпускает град стрел.", a.name)
}
func Fight(characters []Character) {
	for _, c := range characters {

		fmt.Println(c.Attack())
	}
}
func main() {
	warrior, err := NewWarrior("Король Артур")
	if err != nil {
		fmt.Println(err)
		return
	}
	mage, err := NewMage("Мерлин")
	if err != nil {
		fmt.Println(err)
		return
	}
	archer, err := NewArcher("Леголас")
	if err != nil {
		fmt.Println(err)
		return
	}
	characters := []Character{warrior, mage, archer}
	Fight(characters)
}

// https://stepik.org/lesson/1500879/step/5?unit=1520996
