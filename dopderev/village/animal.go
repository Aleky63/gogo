package village

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Animal struct {
	Name   string
	Age    int
	Type   string
	Alive  bool
	Events []string
}

func NewAnimal(name, animalType string, age int) *Animal {
	return &Animal{
		Name:   name,
		Age:    age,
		Type:   animalType,
		Alive:  true,
		Events: []string{},
	}
}

func (a *Animal) IsAlive() bool {
	return a.Alive
}

func (a *Animal) addYear() {
	a.Age++
}

func (a *Animal) die() {
	a.Alive = false
	a.Events = append(a.Events, fmt.Sprintf("aaaaa na %d zzzz.", a.Age))
}

func (a *Animal) Update() {
	if !a.Alive {
		return
	}
	a.addYear()
	if rand.IntN(100) < 5 {
		a.Events = append(a.Events, "Slomal lapy")
	}
	if rand.IntN(100) < 5 && a.Type == "sobaka" {
		a.Events = append(a.Events, "Ukysil soseda")
	}
	if rand.IntN(100) < 15 && a.Type == "koshka" {
		a.Events = append(a.Events, "Ubeshal iz doma")
	}

	if rand.IntN(100) < 8 {
		a.die()
	}
}

func (a *Animal) FlushInfo() string {

	info := fmt.Sprintf("Animal %s ymer v age %d.", a.Name, a.Age)
	if a.Alive {

		events := "net"
		if len(a.Events) > 0 {
			events = strings.Join(a.Events, "\n")
		}
		info = fmt.Sprintf("Animal %s (age: %d),  \nSobaatie: %s\n", a.Name, a.Age, events)
	}
	a.Events = []string{}
	return info
}
