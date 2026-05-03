package main

import (
	"errors"
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Balance int
}

func Pay(user *User, usd int) error {
	if user.Balance-usd < 0 {
		return errors.New("-Недостаточно средств")
	}

	user.Balance -= usd
	return nil
}

func main() {
	user := User{
		Name:    "Putler",
		Balance: 10,
	}
	pp.Println(user)
	err := Pay(&user, 9)
	pp.Println(user)

	if err != nil {
		fmt.Println(" НЕ оплатили", err.Error())
	} else {
		fmt.Println(" оплатили")
	}
}
