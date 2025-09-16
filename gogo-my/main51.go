package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const (
	MinPasswordLength = 4
	MinPasswordsCount = 1
	MaxPasswordsCount = 50
)

var (
	ErrPasswordLengthTooLow = errors.New("password length too low")
	ErrTooLowPasswordsCount = errors.New("too low passwords count")
	ErrTooBigPasswordsCount = errors.New("too big passwords count")
)

var (
	upperChars   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerChars   = []rune("abcdefghijklmnopqrstuvwxyz")
	digitChars   = []rune("0123456789")
	specialChars = []rune("!@#$%^&*")
)

// безопасная перестановка (Fisher–Yates) с crypto/rand
func secureShuffle(runes []rune) {
	for i := len(runes) - 1; i > 0; i-- {
		jBig, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			panic(fmt.Sprintf("crypto/rand shuffle failed: %v", err))
		}
		j := int(jBig.Int64())
		runes[i], runes[j] = runes[j], runes[i]
	}
}

// выбор случайного символа
func getRandomChar(charSet []rune) rune {
	maxIndex := big.NewInt(int64(len(charSet)))
	index, err := rand.Int(rand.Reader, maxIndex)
	if err != nil {
		panic(fmt.Sprintf("crypto/rand failed: %v", err))
	}
	return charSet[index.Int64()]
}

func generateUniquePassword(length int) (string, error) {
	if length < MinPasswordLength {
		return "", ErrPasswordLengthTooLow
	}

	var password strings.Builder

	// гарантируем наличие хотя бы по одному символу из каждого набора
	password.WriteRune(getRandomChar(upperChars))
	password.WriteRune(getRandomChar(lowerChars))
	password.WriteRune(getRandomChar(digitChars))
	password.WriteRune(getRandomChar(specialChars))

	// объединяем все наборы в один
	allChars := append([]rune{}, upperChars...)
	allChars = append(allChars, lowerChars...)
	allChars = append(allChars, digitChars...)
	allChars = append(allChars, specialChars...)

	// заполняем оставшиеся символы
	for i := 4; i < length; i++ {
		password.WriteRune(getRandomChar(allChars))
	}

	// перемешиваем, чтобы первые 4 символа не шли подряд
	runes := []rune(password.String())
	secureShuffle(runes)

	return string(runes), nil
}

func generatePassword(length int, count int) ([]string, error) {
	if length < MinPasswordLength {
		return nil, ErrPasswordLengthTooLow
	}
	if count < MinPasswordsCount {
		return nil, ErrTooLowPasswordsCount
	}
	if count > MaxPasswordsCount {
		return nil, ErrTooBigPasswordsCount
	}

	uniquePasswords := make(map[string]struct{})
	result := make([]string, 0, count)

	for len(result) < count {
		pwd, err := generateUniquePassword(length)
		if err != nil {
			return nil, err
		}
		if _, exists := uniquePasswords[pwd]; !exists {
			uniquePasswords[pwd] = struct{}{}
			result = append(result, pwd)
		}
	}

	return result, nil
}

func main51() {
	passwords, err := generatePassword(19, 9)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Сгенерированные пароли:\n")
	for _, pwd := range passwords {
		fmt.Println(pwd)
	}
}

// https://stepik.org/lesson/1500878/step/2?thread=solutions&unit=1520995
