package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	coffee := "Expresso"

	pointer := &coffee

	fmt.Println("Coffee name:", coffee)
	red := color.New(color.FgRed).SprintfFunc()

	fmt.Printf("Coffee name: %s\n", red("%s", coffee))
	fmt.Println("Memory location:", pointer)
	fmt.Printf("Pointer address: %p\n", pointer)

	fmt.Println("--")

	coffeeCopy := coffee

	fmt.Printf("CoffeeCopy name: %s\n", red("%s", coffeeCopy))
	fmt.Println("Memory location:", &coffeeCopy)
	fmt.Printf("Pointer address: %p\n", &coffeeCopy)
}
