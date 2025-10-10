package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println()
	fmt.Print()
	proba := 5555555
	blue := color.New(color.FgHiBlue).SprintfFunc()
	fmt.Println(blue("%v", proba))
}
