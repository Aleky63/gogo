package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	IsClose     bool
	Rating      float64
}

func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
) (User, error) {
	if name == "" {
		return User{}, errors.New("name cannot be empty")
	}
	if age <= 0 || age >= 150 {
		return User{}, errors.New("age must be between 1 and 149")
	}
	if phoneNumber == "" {
		return User{}, errors.New("phone number cannot be empty")
	}
	if rating <= 0 || rating >= 10 {
		return User{}, errors.New("rating must be greater than 0 and less than 10")
	}

	return User{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNumber,
		IsClose:     isClose,
		Rating:      rating,
	}, nil
}

func (u *User) ChangeName(newName string) {
	if newName != "" {
		u.Name = newName
	}
}

func (u *User) ChangePhone(newPhone string) {
	if newPhone != "" {
		u.PhoneNumber = newPhone
	}
}
func (u *User) ChangeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	}
}
func (u *User) CloseAccount() {
	u.IsClose = true
}
func (u *User) openAccount() {
	u.IsClose = false
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
	}
}
func (u *User) RaingDoun(rating float64) {
	if u.Rating-rating >= 0.0 {
		u.Rating -= rating
	}
}

func main() {
	user, err := NewUser("RRR", 80, "+7(999) 999 9999", true, 5.5)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	fmt.Println("User:", user)
}
