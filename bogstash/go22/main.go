package main

import (
	"fmt"
	// "github.com/fatih/color"
)

func calculatePriceAfterDiscount(price float64, discountRate float64) float64 {
	fmt.Println("Basic coffee price: ", &price)
	return price - (price * discountRate)

}

func main() {

	// red := color.New(color.FgRed).SprintfFunc()
	// fmt.Printf("Coffee name: %s\n", red("%s", coffee))

	var coffeePrice float64 = 5.00
	var discount float64 = 0.1
	fmt.Printf("Basic coffee price: $%.2f\n", coffeePrice)

	coffeePrice = calculatePriceAfterDiscount(coffeePrice, discount)

	fmt.Printf("Basic coffee price: $%.2f\n", coffeePrice)
	fmt.Println("Basic coffee price: ", &coffeePrice)

}
