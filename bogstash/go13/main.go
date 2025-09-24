package main

import (
	"fmt"

	"github.com/fatih/color"
)

func greet(coffeeShopName string) {
	fmt.Println("fefefefefefefefefef", coffeeShopName)
}
func calc(amountSpent float32) int {
	loyaltyPoints := int(amountSpent * 20)

	return loyaltyPoints
}

func main() {
	greet("555")
	greet("777")
	color.Blue("QCQCWDSFDVVFVVS")
	yyyy := 555555555
	magenta := color.New(color.FgMagenta).PrintfFunc()
	magenta("Is my: %v", yyyy)
	fmt.Println()
	var newCalc int = calc(6.2)
	color.Red("%d", newCalc)
	fmt.Print(newCalc)

}
