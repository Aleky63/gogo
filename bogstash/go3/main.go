package main

import "fmt"

func main() {
	name := "AAAAAAAAAAAA"
	var price float32 = 2.99
	ready := true
	count := 5

	fmt.Printf("Type is: \n%s\n%T\n%T\n%T\n%T", name, name, price, ready, count)

}
