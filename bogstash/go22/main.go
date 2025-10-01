package main

import (
	"fmt"

	"github.com/fatih/color"
)

func applyDiscount(price *float64, discountRate float64) {

	*price = *price - (*price * discountRate)

}

func main() {

	yellow := color.New(color.FgHiYellow).SprintfFunc()
	// fmt.Printf("Coffee name: %s\n", red("%s", coffeePrice))

	var coffeePrice float64 = 5.00
	var discount float64 = 0.1
	fmt.Printf("Basic coffee price: $%s\n", yellow("%.2f", coffeePrice))

	applyDiscount(&coffeePrice, discount)

	fmt.Printf("New coffee price: $%s\n", yellow("%.2f", coffeePrice))

}
