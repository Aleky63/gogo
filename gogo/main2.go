package main

import (
	"fmt"
	"log"
	"math"
)

func main2() {

var weight, height float64

fmt.Printf("Введите ваш вес (кг): ")

if _ , err := fmt.Scan(&weight); err != nil {
	log.Fatalf("Ошибка ввода веса: %s\n", err)
}


fmt.Printf("Введите ваш рост (см): ")

	if _ , err := fmt.Scan(&height ); err != nil {
	log.Fatalf("Ошибка ввода роста: %s\n", err)
}
	

bmi := weight / math.Pow(height/100, 2)



var category string
switch {
	case bmi < 18.5:
		category = "Недостаточный вес"
	case bmi < 25:
		category = "Нормальный вес"
	case bmi < 30:
		category = "Избыточный вес"
	default:
		category = "Ожирение"
	}
fmt.Printf("Ваш ИМТ: %.2f\n", bmi)
	fmt.Printf("Категория: %s\n", category)

}