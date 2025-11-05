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
		err := errors.New("nедостаточно средств!")
		return err
	}
	user.Balance -= usd

	return nil

}

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("panika:", p)
		}
	}()

	user := User{
		Name:    "Donny",
		Balance: 50,
	}
	pp.Println(user)
	str := Pay(&user, 12)

	if str == nil {
		fmt.Println("Оплата прошла")
	} else {
		fmt.Println("Оплата  не прошла, error: ", str)
	}
	pp.Println(user)

}
