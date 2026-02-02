package main

import (
	"animals/pets"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
red := color.New(color.FgRed).SprintFunc()
	myCat := pets.Cat{Name: "mr.Buttons"}
	myDog := pets.Dog{Name: "Tramp"}

fmt.Println(red("Мой кот:", myCat.Name))
fmt.Println(red("Мой пес:", myDog.Name))


var feedToCat int = 2
var feedToDog int = 6

catFed, err := feed (&myCat, feedToCat)

if err != nil {
	fmt.Fprintf(os.Stderr, "Error feeding cat: %v\n", err)
 } else {
	fmt.Println("Cat ate:", catFed)
}


fmt.Println(red("...---------////////---------...."))



dogFed, err := feed (&myDog, feedToDog)

if err != nil {
	fmt.Fprintf(os.Stderr, "Error feeding dog: %v\n", err)
 } else {
	fmt.Println("Dog ate:", dogFed)
}

}

func feed (animal pets.EaterWalker, amount int)(int, error){
	fmt.Println("First, let's walk!!!")
	fmt.Println(animal.Walk())
	
return  animal.Eat(amount)
}