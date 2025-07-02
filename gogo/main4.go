package main

import (
	"fmt"
	"os"
)

func main() {

var input int
fmt.Println("Введите время в часах (0-23): ")
_, err := fmt.Scan(&input)


if  err != nil || input < 0 || input > 23 {
fmt.Println( "Неверно задано время.")
os.Exit(1)
}

fmt.Println(input)


var x string
switch {
	case input > 23 && input < 6 :
	x = "Ночь "
	case input < 12 :
	x = "Утро " 
	case input < 18 :
		x = "День "
	default:
		x = "Вечер "
	}

fmt.Printf("Сейчас %02dч. - %s\n", input, x)
}



// 	log.Fatalf("Ошибка ввода веса: %s\n", err)
// }
// if _ , err := fmt.Scan(&weight); err != nil {
// 	log.Fatalf("Ошибка ввода веса: %s\n", err)
// }


// fmt.Printf("Введите ваш рост (см): ")

// 	if _ , err := fmt.Scan(&height ); err != nil {
// 	log.Fatalf("Ошибка ввода роста: %s\n", err)
// }
	

// bmi := weight / math.Pow(height/100, 2)






// var category string
// switch {
// 	case bmi < 18.5:
// 		category = "Недостаточный вес"
// 	case bmi < 25:
// 		category = "Нормальный вес"
// 	case bmi < 30:
// 		category = "Избыточный вес"
// 	default:
// 		category = "Ожирение"
// 	}
// fmt.Printf("Ваш ИМТ: %.2f\n", bmi)
// 	fmt.Printf("Категория: %s\n", category)

