package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type VillageElement interface {
	Update()
	FlushInfo() string
}

type Resident struct {
	Name    string
	Age     int
	Married bool
	Alive   bool
	Events  []string
}

func (r *Resident) addYear() {
	r.Age++
}
func (r *Resident) changeMarriedStatus() {
	if r.Married {
		r.Married = false
		r.Events = append(r.Events, "Razvod")
	} else {
		r.Married = true
		r.Events = append(r.Events, "Zazzyzy")
	}
}

func (r *Resident) die() {
	r.Alive = false
	r.Events = append(r.Events, fmt.Sprintf("www na %d zzzz.", r.Age))
}

func (r *Resident) Update() {
	if !r.Alive {
		return
	}
	r.addYear()
	if rand.IntN(100) < 15 {
		r.changeMarriedStatus()
	}
	if rand.IntN(100) < 15 {
		r.Events = append(r.Events, "Razvelcy")
	}
	if rand.IntN(100) < 15 {
		r.Events = append(r.Events, "Nashel work")
	}
	if r.Married && rand.IntN(100) < 25 {
		r.Events = append(r.Events, "Porygalsy")
	}
	if rand.IntN(100) < 3 {
		r.die()
	}
}

func (r *Resident) FlushInfo() string {

	info := fmt.Sprintf("Residenttt %s ymer v age %d.", r.Name, r.Age)
	if r.Alive {
		marriedStatus := "holost"

		if r.Married {
			marriedStatus = "v brake"
		}
		events := "net"
		if len(r.Events) > 0 {
			events = strings.Join(r.Events, "\n")
		}
		info = fmt.Sprintf("Residenttt %s (age: %d), status: %s. \nSobaatie: %s\n", r.Name, r.Age, marriedStatus, events)
	}
	r.Events = []string{}
	return info
}

type Animal struct {
	Name   string
	Age    int
	Type   string
	Alive  bool
	Events []string
}

type Village struct {
	Elements []VillageElement
}

func (v *Village) AddElement(e VillageElement) {
	v.Elements = append(v.Elements, e)
}

func (v *Village) UpdateAll() {
	for _, e := range v.Elements {
		e.Update()
	}
}

func (v Village) ShowAllInfo() string {
	info := ""
	for _, e := range v.Elements {
		info += e.FlushInfo()
	}
	return info
}

func main() {
	village := Village{}

	// Создаем жителей деревни
	resident1 := &Resident{Name: "Алиса", Age: 30, Married: false, Alive: true, Events: []string{}}
	resident2 := &Resident{Name: "Борис", Age: 40, Married: true, Alive: true, Events: []string{}}

	// Создаем животных
	// animal1 := &Animal{Name: "Бобик", Age: 5, Type: "собака", Alive: true, Events: []string{}}
	// animal2 := &Animal{Name: "Мурка", Age: 3, Type: "кошка", Alive: true, Events: []string{}}

	// Добавляем элементы в деревню
	village.AddElement(resident1)
	village.AddElement(resident2)
	// village.AddElement(animal1)
	// village.AddElement(animal2)

	// Симуляция обновления деревни на несколько лет
	for i := 0; i < 5; i++ {
		fmt.Printf("Год %d:\n", i+1)
		village.UpdateAll()
		fmt.Println(village.ShowAllInfo())
	}
}

// https://stepik.org/lesson/1538383/step/2?unit=1558981
