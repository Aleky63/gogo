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

func processPayment(orderTotal float64, tip float64, amountPaid float64) (totalAmountDue float64, change float64) {
	totalAmountDue = orderTotal + tip
	change = amountPaid - totalAmountDue
	return
}

func main() {
	greet("555")

	yyyy := 555555555
	magenta := color.New(color.FgMagenta).PrintfFunc()
	magenta("Is my: %v", yyyy)
	fmt.Println("")
	var newCalc int = calc(6.2)
	color.Red("%d", newCalc)
	fmt.Println(newCalc)

	tot, ch := processPayment(216.00, 20.00, 500.00)
	fmt.Println(tot, ch)
}
