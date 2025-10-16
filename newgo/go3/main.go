package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {

	fmt.Print("A  ")
	time.Sleep(3 * time.Second)
	fmt.Println(aboba(5, 6, 9))
}
func aboba(a, b, c int) string {
	red := color.New(color.FgHiRed).SprintFunc()

	return "Sum: " + red(a+b+c)

}
