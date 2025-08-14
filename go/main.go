package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type User struct {
	FirstName         string
	LastName          string
	BirthYear         int
	FavoriteLanguages []string
}

func (u User) SecretIdentity() (string, error) {
	if u.FirstName == "" || u.LastName == "" {
		return "", errors.New("имя или фамилия пустые")
	}

	firstNameInitial := strings.ToUpper(string([]rune(u.FirstName)[0]))
	lastNameInitial := strings.ToUpper(string([]rune(u.LastName)[0]))

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1

	secretID := fmt.Sprintf("%s%s%d", firstNameInitial, lastNameInitial, randomNumber)
	return secretID, nil
}


func main() {
	
	user := User{
		FirstName: "Алексей",
		LastName:  "Смирнов",
	}
	
	fmt.Printf("Пользователь: %s %s\n", user.FirstName, user.LastName)

	if secret, err := user.SecretIdentity(); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Секретное имя: %s\n", secret)
	}
}