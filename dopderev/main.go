package main

import (
	"dopderev/village"
	"fmt"
)

func main() {
	v := village.NewVillage()

	// Создаем жителей деревни
	resident1 := village.NewResident("Алиса", 30, false)
	resident2 := village.NewResident("Борис", 40, true)
	// Создаем животных
	animal1 := village.NewAnimal("Бобик", "собака", 5)
	animal2 := village.NewAnimal("Мурка", "кошка", 3)
	// Добавляем элементы в деревню
	v.AddElement(resident1)
	v.AddElement(resident2)
	v.AddElement(animal1)
	v.AddElement(animal2)

	// Симуляция обновления деревни на несколько лет
	for i := 0; i < 5; i++ {
		fmt.Printf("Год %d:\n", i+1)
		v.UpdateAll()
		fmt.Println(v.ShowAllInfo())

	}
}
