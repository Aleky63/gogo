package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	var price float64
	name := "AAAAAAAAAAAA"
	price = 3.99
	ready := true
	count := 5

	var (
		customerName string = "Bogdan"
		tableNumber  int    = 8
		isReadyToPay bool   = true
	)

	fmt.Printf("Type is: \n%s\n%T\n%T\n%T\n%T\n\n", name, name, price, ready, count)

	name = "BBBBBBBBBBBBB"

	magenta := color.New(color.FgMagenta).PrintfFunc()

	magenta("Type is new: \n%s\n%T\n%T\n%T\n%T\n ", name, name, price, ready, count)

	red := color.New(color.FgRed).PrintfFunc()
	red("\nPrice of the coffee is: %.2f\n", price)

	total := price * float64(count)

	fmt.Println(total)

	const weekRewardPoints = 10
	fmt.Printf("SUM %T\n", weekRewardPoints)

	var totalRewardPoints float64 = 302.3
	totalRewardPoints = totalRewardPoints + weekRewardPoints
	fmt.Println(totalRewardPoints)
	fmt.Println(customerName, tableNumber, isReadyToPay)
	fmt.Printf("Customer %s is %d %t", customerName, tableNumber, isReadyToPay)
}
