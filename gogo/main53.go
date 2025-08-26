package main

import "fmt"

type Dog struct{}
type Cat struct{}
type Bird struct{}

func (Dog) Speak() string {
	return "Ggggggggg"
}
func (Dog) Growl() string {
	return "Rrrrrrrr"
}
func (Cat) Speak() string {
	return "Mmmmmmy"
}
func (Bird) Speak() string {
	return "Kyyyyy"
}

func main() {
	dog := Dog{}
	cat := Cat{}
	bird := Bird{}

	makeSound(dog)
	makeSound(cat)
	makeSound(bird)
}

type Speaker interface {
	Speak() string
}

func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}
