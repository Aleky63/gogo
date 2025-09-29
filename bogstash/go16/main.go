package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	texRate := 0.05
	calculateTax := func(amout float64) float64 {
		return amout * texRate
	}
	subtotal := 25.00
	tax := calculateTax(subtotal)
	total := subtotal + tax

	fmt.Println(total)
	red := color.New(color.FgRed).PrintfFunc()
	red("res: %.2f\n", total)

}
