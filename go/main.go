package main

import (
	"errors"
	"fmt"
	"log"
)

type Engine struct {
	HorsePower int
	Started    bool
}

func (e *Engine) Start() error {
	if e.Started {
		return errors.New(("already started"))
	}
	e.Started = true
	return nil
}

type Car struct {
	Engine Engine
	Model  string
}

func (c *Car) Start() error {
	if err := c.Engine.Start(); err != nil {
		return fmt.Errorf("engine start: %w", err)
	}
	return nil
}

func (c Car) Drive() error {
	if !c.Engine.Started {
		return errors.New("engine not started")
	}
	return nil
}

func main() {
	car := Car{
		Engine: Engine{
			Started:    false,
			HorsePower: 150,
		},
		Model: "Toyota",
	}

	if err := car.Drive(); err != nil {
		log.Fatalf("drive car: %v", err)
	}

	fmt.Printf("auto pusk, model %s, sil: %d", car.Model, car.Engine.HorsePower)

	fmt.Printf("%+v\n", car)

	if err := car.Start(); err != nil {
		log.Fatalf("start car: %v", err)
	}

}
