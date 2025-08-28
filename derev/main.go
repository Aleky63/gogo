package main

import "fmt"

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
	animal1 := &Animal{Name: "Бобик", Age: 5, Type: "собака", Alive: true, Events: []string{}}
	animal2 := &Animal{Name: "Мурка", Age: 3, Type: "кошка", Alive: true, Events: []string{}}

	// Добавляем элементы в деревню
	village.AddElement(resident1)
	village.AddElement(resident2)
	village.AddElement(animal1)
	village.AddElement(animal2)

	// Симуляция обновления деревни на несколько лет
	for i := 0; i < 5; i++ {
		fmt.Printf("Год %d:\n", i+1)
		village.UpdateAll()
		fmt.Println(village.ShowAllInfo())
	}
}
