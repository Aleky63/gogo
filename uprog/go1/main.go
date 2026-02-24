package main

import (
	"fmt"

	"github.com/fatih/color"
)



func greet (name string){
	fmt.Println(name,"!!!")
}


func main() {
greet("Tramp")

var (
red = color.New(color.FgHiRed).SprintFunc()
blue = color.New(color.FgHiBlue).SprintFunc()

)

	fmt.Println (red(555) + blue (" dfdfdfdefd  "))
}


