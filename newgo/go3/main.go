package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {

	fmt.Print("A  ")
	time.Sleep(3 * time.Second)
	fmt.Println(aboba(5, 6, 8))
}
func aboba(a, b, c int) string {
	red := color.New(color.FgHiRed).SprintFunc()
	sum := a + b + c
	return "Sum: " + red(sum)

}
