package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name     string
	Ballance int
}

func Pay(user *User, usd int) bool {
	if user.Ballance-usd >= 0 {
		user.Ballance -= usd
		return true
	}
	return false
}

func main() {
	user := User{
		Name:     "Putler",
		Ballance: 10,
	}
	pp.Println(user)
	payed := Pay(&user, 16)
	pp.Println(user)

	if payed {
		fmt.Println("оплатили")
	} else {
		fmt.Println(" НЕ оплатили")
	}
}
