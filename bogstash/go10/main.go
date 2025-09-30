package main

import (
	"fmt"

	"github.com/fatih/color"
)

func printDrinkInfo(name string, drink string, price float64) {
	red := color.New(color.FgRed).SprintFunc()
	info := "%s's favorite drink is %s and it's price is %s\n	"

	fmt.Printf(info, name, drink, red(fmt.Sprintf("$%.2f", price)))

}
func main() {
	printDrinkInfo("Tramp", "Chay", 4.66)

}
