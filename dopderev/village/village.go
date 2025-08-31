package village

import (
	"fmt"
	"math/rand/v2"
)

type VillageElement interface {
	Update()
	FlushInfo() string
	IsAlive() bool
}

type Village struct {
	Elements []VillageElement
	year     int
}

func NewVillage() *Village {
	return &Village{
		Elements: []VillageElement{},
	}
}

func (v *Village) AddElement(e VillageElement) {
	v.Elements = append(v.Elements, e)
}

func (v *Village) UpdateAll() {
	v.year++

	newElements := []VillageElement{}

	for _, e := range v.Elements {
		if e.IsAlive() {
			e.Update()
		}
	}

	// шанс рождения детей у жителей
	for _, e := range v.Elements {
		if r, ok := e.(*Resident); ok && r.Alive && r.Married && r.Age < 50 {
			if rand.IntN(100) < 10 { // 10% шанс
				childName := fmt.Sprintf("Ребёнок_%d", rand.IntN(1000))
				child := NewResident(childName, 0, false)
				newElements = append(newElements, child)
				r.Events = append(r.Events, fmt.Sprintf("У них родился ребёнок: %s", childName))
			}
		}
	}

	// шанс появления новых животных
	for _, e := range v.Elements {
		if a, ok := e.(*Animal); ok && a.Alive {
			if rand.IntN(100) < 8 { // 8% шанс
				babyName := fmt.Sprintf("%s_детёныш", a.Name)
				baby := NewAnimal(babyName, a.Type, 0)
				newElements = append(newElements, baby)
				a.Events = append(a.Events, fmt.Sprintf("У животного родился детёныш: %s", babyName))
			}
		}
	}

	// добавляем новорожденных в деревню
	v.Elements = append(v.Elements, newElements...)
}

func (v Village) ShowAllInfo() string {
	info := fmt.Sprintf("Год: %d\n", v.year)
	for _, e := range v.Elements {
		info += e.FlushInfo()
	}
	return info
}
