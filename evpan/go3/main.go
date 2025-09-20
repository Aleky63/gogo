package main

import (
	"fmt"
	u "go3/utils"
)

func main() {

	alice := u.Person{Name: "Alisssse", Age: 105}
	fmt.Println(alice.Age, alice.Name)
	u.Greet(alice.Name)
}
