package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {


var (
red = color.New(color.FgHiRed).SprintFunc()
blue = color.New(color.FgHiBlue).SprintFunc()

)

	fmt.Println (red(555) + blue (" dfdfdfdefd  "))
}


