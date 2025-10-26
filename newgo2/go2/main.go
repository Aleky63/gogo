package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println()
	fmt.Print()
	proba := "5555ffff555"
	blue := color.New(color.BgHiBlue).SprintfFunc()
	fmt.Println(blue("%v", proba))
}
