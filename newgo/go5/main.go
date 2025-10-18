package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	IsClose     bool
	Rating      float64
}

func aboba(a, b, c int) string {

	red := color.New(color.FgHiRed).SprintFunc()
	sum := a + b + c
	return "Sum: " + red(sum)

}

func (u User) Greeting() {
	fmt.Println("Meyz zovyt:", u.Name)
	fmt.Println("Meyz age:", u.Age)
	fmt.Println("Meyz isclose:", u.IsClose)
	fmt.Println("Meyz tel:", u.PhoneNumber)
	fmt.Println("Meyz rat:", u.Rating)
}
func (u *User) RaingUp(rating float64) {
	if u.Rating+rating <= 10.0 {
		u.Rating += rating
		fmt.Println(" прошел")
	} else {
		fmt.Println(("Не прошел"))
	}
}

func main() {
	defer func() {
		fmt.Print("После всех вывести ")
	}()
	user := User{
		Name:        "RRR",
		Age:         100,
		PhoneNumber: "+7(999) 999 9999",
		IsClose:     false,
		Rating:      5.5,
	}
	fmt.Println("User:", user)

	fmt.Println("A")
	time.Sleep(3 * time.Second)
	fmt.Println(aboba(5, 6, 8))

	user.Greeting()
	user.RaingUp(3.33)
	fmt.Println(user)
}
