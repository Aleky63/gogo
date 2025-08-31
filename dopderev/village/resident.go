package village

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Resident struct {
	Name    string
	Age     int
	Married bool
	Alive   bool
	Events  []string
}

func NewResident(name string, age int, married bool) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Married: married,
		Alive:   true,
		Events:  []string{},
	}
}

func (r *Resident) IsAlive() bool {
	return r.Alive
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
