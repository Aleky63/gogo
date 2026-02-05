package main

import (
	"animals/pets"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
red := color.New(color.FgRed).SprintFunc()
	myCat := pets.Cat{

	Animal: pets.Animal {Name: "mr.Buttons"},
	Age: 7,
	IsAsleep: true,
	}
	myDog := pets.Dog{
		 Animal: pets.Animal {Name: "Tramp"},
		Age: 8,
		Weight: 15,
	}

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
displayInfo(myCat)
displayInfo(myDog)
displayInfo(66)
}

func feed (animal pets.EaterWalker, amount int)(int, error){
	fmt.Println("First, let's walk!!!")
	fmt.Println(animal.Walk())
	fmt.Println("Now,let's feed our", animal.GetName())
	
return  animal.Eat(amount)
}


func displayInfo(i interface{}){
	switch v :=  i.(type){
	case pets.Cat:
		fmt.Println("This is a Cat named:",v.Name,"and it is",v.Age,"years old")
	case pets.Dog:
		fmt.Println("This is a Dog named:",v.Name,"and it is",v.Age,"years old","and weight" ,v.Weight, "kg")
	default:
		fmt.Println("This is unknowm")
	}
}