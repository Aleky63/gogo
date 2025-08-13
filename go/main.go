package main

import "fmt"
type Name struct{
	First string
	Last string
}

type User struct {
	Name Name
	Year int
}


func(u User) Greet() {

fmt.Printf("Z %s %s %d   fdfddfdfdf.\n", u.Name.First, u.Name.Last, u.Year)
}

func main() {
	user := User {
		Name: Name{
			First:"AAAAAA",
	        Last : "dddddd",
		},
		Year: 19999999,
	}
	user2 := User {
		Name: Name{
			First:"OOOOOO",
	        Last : "aa",
		},
		Year: 99б
	}

	fmt.Printf("%+v\n", user)
	user.Greet()
	user.Name.First = "YYYYYYYYYYYYYYYYYYYYYYYYYYYYY"
	user.Greet()
	user2.Greet()
}





	

