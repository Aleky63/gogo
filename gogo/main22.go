package main

import (
	"fmt"
	"strings"
)

func CountVowelsInArray(arr [3]string) {
    vowels := "аеёиоуыэюяАЕЁИОУЫЭЮЯ"
    results := make([]int, 3)

    for i, str := range arr {
        count := 0
        for _, char := range str {
            if strings.ContainsRune(vowels, char) {
                count++
            }
        }
        results[i] = count // Присваиваем результат
    }

    fmt.Printf("%d %d %d\n", results[0], results[1], results[2])
}

func main22() {
    arr := [3]string{"Привет", "Счастье", "Мир"}
    CountVowelsInArray(arr)
}
// func CountVowelsInArray(arr [3]string) {
// 	for i := 0; i < len(arr); i++ {
// 		count := 0
// 		for _, char := range arr[i] {
// 			switch char {
// 			case 'а', 'А', 'е', 'Е', 'ё', 'Ё', 'и', 'И', 'о', 'О',
// 				'у', 'У', 'ы', 'Ы', 'э', 'Э', 'ю', 'Ю', 'я', 'Я':
// 				count++
// 			}
// 		}
// 		fmt.Print(count, " ")
// 	}
// }