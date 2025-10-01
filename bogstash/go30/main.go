package main

import (
	"fmt"

	"github.com/fatih/color"
)

//	func main() {
//		coffeeTypes := [3]string{"Tramp", "Putler", "Si"}
//		yellow := color.New(color.FgHiYellow).SprintfFunc()
//		fmt.Println(coffeeTypes)
//		fmt.Println(yellow("%v", coffeeTypes))
//	}
func main() {

	yellow := color.New(color.FgHiYellow).SprintfFunc()
	coffeeTypes := [3]string{"Tramp", "Putler", "Si"}

	slice := coffeeTypes[1:]

	fmt.Println(coffeeTypes)
	fmt.Println(yellow("%v", slice))
}
