package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	coffeName := "Expresso"
	size := "Medium"
	price := 2.99

	color.Blue("QCQCWDSFDVVFVVS")

	fmt.Println("Medium Expresso price is $3.00")

	color.Red("Medium Expresso price is $3.00")

	fmt.Println(size, coffeName, "price is $", price)

	fmt.Printf("%v %v price is $%.2f\n", size, coffeName, price)

	magenta := color.New(color.FgMagenta).PrintfFunc()
	magenta("%s %v price is $%.2f\n", size, coffeName, price)
}
