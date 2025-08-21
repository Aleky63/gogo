package main

import (
	"fmt"
)

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Printf("%s издает звук\n", a.Name)
}

type Dog struct {
	Animal
}

func (d Dog) Speak() {
	fmt.Printf("%s издает звук\n", d.Name)
}
func main() {
	dog := Dog{

		Animal: Animal{
			Name: "EWEDAcdcd",
		},
	}
	dog.Speak()
}
