package main

import (
	"fmt"

	"github.com/fatih/color"
)

func getDrinkInfo(name string, drink string, price float64) {
	red := color.New(color.FgRed).SprintFunc()
	info := "%s's favorite drink is %s and its price is %s\n	"

	fmt.Printf(info, name, drink, red(fmt.Sprintf("$%.2f", price)))

}
func main() {
	getDrinkInfo("Tramp", "Chay", 4.66)

}
