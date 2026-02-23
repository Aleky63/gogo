package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
red := color.New(color.FgHiRed).SprintFunc()
	fmt.Println (red(555, "dfdfdfdfdefd"))
}