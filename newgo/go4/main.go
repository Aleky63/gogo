package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {
	defer func() {
		fmt.Print("C")
	}()
	fmt.Println("A")
	time.Sleep(3 * time.Second)
	fmt.Println(aboba(5, 6, 8))

	number := 11111111
	foo(&number)
	fmt.Println(number)
}
func aboba(a, b, c int) string {

	red := color.New(color.FgHiRed).SprintFunc()
	sum := a + b + c
	return "Sum: " + red(sum)

}
func foo(n *int) {

	fmt.Println(n)
	fmt.Println(*n)
	*n = 5555
}
