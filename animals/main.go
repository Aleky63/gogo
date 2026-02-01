package main

import (
	"animals/pets"
	"fmt"

	"github.com/fatih/color"
)

func main() {
red := color.New(color.FgRed).SprintFunc()
	myCat := pets.Cat{Name: "mr.Buttons"}
	// myDog := pets.Dog{Name: "spot"}

fmt.Println(red("Мой кот:", myCat.Name))
}

func feed (amount uint8)(uint8, error){

}