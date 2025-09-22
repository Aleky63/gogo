package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	name := "AAAAAAAAAAAA"
	var price float32 = 2.99
	ready := true
	count := 5

	fmt.Printf("Type is: \n%s\n%T\n%T\n%T\n%T\n\n", name, name, price, ready, count)

	name = "BBBBBBBBBBBBB"

	magenta := color.New(color.FgMagenta).PrintfFunc()

	magenta("Type is new: \n%s\n%T\n%T\n%T\n%T\n", name, name, price, ready, count)
}
