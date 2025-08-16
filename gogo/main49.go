package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

type User struct {
	FirstName         string
	LastName          string
	BirthYear         int
	FavoriteLanguages []string
}

func (u User) SecretIdentity() string {
	if len(u.FirstName) == 0 || len(u.LastName) == 0 {
		return ""
	}

	firstLetter := string([]rune(u.FirstName)[0])
	lastLetter := string([]rune(u.LastName)[0])
	randomNum := rand.Intn(100) + 1

	return fmt.Sprintf("%s%s%d", firstLetter, lastLetter, randomNum)
}

func (u User) Age() int {
	return time.Now().Year() - u.BirthYear
}

func (u *User) AddFavoriteLanguage(language string) error {
	if language == "" {
		return errors.New("empty language name")
	}

	for _, lang := range u.FavoriteLanguages {
		if lang == language {
			return errors.New("duplicate")
		}
	}

	u.FavoriteLanguages = append(u.FavoriteLanguages, language)
	return nil
}

func (u *User) RemoveFavoriteLanguage(language string) error {
	for i, lang := range u.FavoriteLanguages {
		if lang == language {
			u.FavoriteLanguages = append(u.FavoriteLanguages[:i], u.FavoriteLanguages[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func (u User) IsProgrammingLanguageFavorite(language string) bool {
	for _, lang := range u.FavoriteLanguages {
		if lang == language {
			return true
		}
	}
	return false
}

func (u User) RandomFavoriteLanguage() (string, error) {
	if len(u.FavoriteLanguages) == 0 {
		return "", errors.New("no options")
	}

	randomIndex := rand.Intn(len(u.FavoriteLanguages))
	return u.FavoriteLanguages[randomIndex], nil
}

func (u User) GenerateProfile() string {
	languages := ""
	if len(u.FavoriteLanguages) > 0 {
		languages = strings.Join(u.FavoriteLanguages, ", ")
	}

	return fmt.Sprintf("Имя: %s.\nФамилия: %s.\nВозраст: %d.\nСписок любимых языков программирования: [%s].",
		u.FirstName, u.LastName, u.Age(), languages)
}

func (u *User) UpdateName(firstName, lastName string) error {
	if firstName == "" || lastName == "" {
		return errors.New("empty data")
	}

	if len(firstName) > 0 && len(lastName) > 0 {
		firstChar := []rune(firstName)[0]
		lastChar := []rune(lastName)[0]

		if !unicode.IsUpper(firstChar) || !unicode.IsUpper(lastChar) {
			return errors.New("invalid data")
		}
	}

	u.FirstName = firstName
	u.LastName = lastName
	return nil
}
func main49() {
	// Создаем пользователя
	user := &User{
		FirstName: "Donald",
		LastName:  "Putlin",
		BirthYear: 1990,
	}

	// Добавляем любимые языки
	if err := user.AddFavoriteLanguage("Go"); err != nil {
		fmt.Println("Ошибка:", err)
	}
	if err := user.AddFavoriteLanguage("Python"); err != nil {
		fmt.Println("Ошибка:", err)
	}
	if err := user.AddFavoriteLanguage("Go"); err != nil { // Дублируем
		fmt.Println("Ошибка:", err)
	}
	if err := user.AddFavoriteLanguage(""); err != nil { // Пустое имя
		fmt.Println("Ошибка:", err)
	}

	// Проверяем, любим ли язык
	fmt.Println("Go любим?:", user.IsProgrammingLanguageFavorite("Go"))
	fmt.Println("Java любим?:", user.IsProgrammingLanguageFavorite("Java"))

	// Удаляем язык
	if err := user.RemoveFavoriteLanguage("Python"); err != nil {
		fmt.Println("Ошибка:", err)
	}
	if err := user.RemoveFavoriteLanguage("Rust"); err != nil { // Нет такого
		fmt.Println("Ошибка:", err)
	}

	// Случайный любимый язык
	lang, err := user.RandomFavoriteLanguage()
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Случайный любимый язык:", lang)
	}

	// Секретное имя
	fmt.Println("Секретное имя:", user.SecretIdentity())

	// Обновление имени
	if err := user.UpdateName("Павел", "Тарасов"); err != nil {
		fmt.Println("Ошибка при обновлении имени:", err)
	}

	// Профиль
	fmt.Println("\nПолный профиль пользователя:")
	fmt.Println(user.GenerateProfile())
}

// https://stepik.org/lesson/1500870/step/5?thread=solutions&unit=1520987
