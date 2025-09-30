package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {

	var coffeePrice = 4.50
	fmt.Println("Coffee price", coffeePrice)
	coffeePrice = 5.55
	fmt.Println("Coffee priceeeee", coffeePrice)
	fmt.Println("Coffee price", &coffeePrice)
	fmt.Println("Coffee priceeeee", &coffeePrice)
	var pointer *float64 = &coffeePrice
	red := color.New(color.FgRed).SprintfFunc()
	*pointer = 6.66
	color := red("%.2f", *pointer)
	fmt.Println(pointer)
	fmt.Println("Coffee priceeeee", color)
}
