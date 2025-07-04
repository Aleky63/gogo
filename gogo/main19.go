// package main

// import (
// 	"bufio"
// 	"errors"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// 	"unicode"
// )
// //
// func GetInput() (str string, err error) {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()

//     if err = scanner.Err(); err != nil {
// 		return "", fmt.Errorf("scanner.Err(): %w", err)
// 	}

// 	str = scanner.Text()

// 	if str == "" {
// 		return "", errors.New("String is empty")
// 	}
// 	if strings.TrimSpace(str) == "" {
// 		return "", errors.New("String contains only spaces")
// 	}
// 	return
// }

// func CountCharacters(text string) (letters, digits, spaces, punctuation int){

// for _, c := range text {
// 		switch {
// 		case c == ' ':
// 			spaces++
// 		case unicode.IsLetter(c):
// 			letters++
// 		case unicode.IsDigit(c):
// 			digits++
// 		case unicode.IsPunct(c):
// 			punctuation++
// 		}
// 	}
// 	return
// }
// func DisplayResults(letters, digits, spaces, punctuation int){
// fmt.Printf("Letters count: %d\n", letters)
// 	fmt.Printf("Digits count: %d\n", digits)
// 	fmt.Printf("Spaces count: %d\n", spaces)
// 	fmt.Printf("Punctuation count: %d\n", punctuation)
//   }
// func main(){
//     text, err := GetInput()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	DisplayResults(CountCharacters(text))
//     fmt.Print("Please enter text: ")
// }

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"unicode"
)

// Ввод текста
// Создайте функцию GetInput() (string, error), которая запрашивает у пользователя ввод текста и возвращает его для дальнейшего анализа. Если функция вернет ошибку, программу необходимо завершить.


func GetInput() (str string, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите текст:")
	 if !scanner.Scan(){
         fmt.Println()
	 
    if err = scanner.Err(); err != nil {
		return "", fmt.Errorf("scanner.Err(): %w", err)
	}
return "", errors.New("unable to read input")
}
return scanner.Text(), nil
}
	
// Подсчет символов
// Реализуйте функцию CountCharacters(text string) (letters, digits, spaces, punctuation int), которая принимает текст и возвращает количество:

// букв (letters)
// цифр (digits)
// пробелов (spaces)
// знаков препинания (punctuation)

func CountCharacters(text string) (letSmalls, letBigs, digits, spaces, punctuation int){

for _, c := range text {
		switch {
		
		case unicode.IsUpper(c):
			letBigs++
		case unicode.IsLower(c) :
			letSmalls++
		case unicode.IsDigit(c):
			digits++
		case unicode.IsPunct(c):
			punctuation++
		case unicode.IsSpace(c):
			spaces++		
		}
	}
	return
}

// Вывод результатов

// Напишите функцию DisplayResults(letters, digits, spaces, punctuation int), которая принимает результаты анализа и выводит их на экран в удобочитаемом формате. Например:

// Количество букв: 50
// Количество цифр: 10
// Количество пробелов: 5
// Количество знаков препинания: 3

func DisplayResults( letBigs, letSmalls, digits, spaces, punctuation int){

// fmt.Printf(`Количество заглавных букв: %d
// Количество прописных букв: %d
// Количество цифр: %d
// Количество пробелов: %d
// Количество знаков препинания: %d`, letSmalls, letBigs, digits, spaces, punctuation )

    fmt.Printf("letBigs count: %d\n", letBigs)
    fmt.Printf("letSmalls count: %d\n", letSmalls)
	fmt.Printf("Digits count: %d\n", digits)
	fmt.Printf("Spaces count: %d\n", spaces)
	fmt.Printf("Punctuation count: %d\n", punctuation)
  }
func main(){
    text, err := GetInput()

	if err != nil {
		log.Fatal(err)
	}

	DisplayResults(CountCharacters(text))
    fmt.Print("Please enter text: ")
}