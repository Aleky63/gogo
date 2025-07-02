// package main

// import (
//     "bufio"
//     "fmt"
//     "math/rand"
//     "os"
//     "strconv"
//     "strings"
//     "time"
// )

// // Генерация случайного числа от min до max (включительно)
// func getRandomNumber(min, max int) int {
//     return rand.Intn(max-min+1) + min
// }

// // Получение ввода от пользователя
// func getUserInput(scanner *bufio.Scanner) string {
//     scanner.Scan()
//     return scanner.Text()
// }

// // Основная функция игры
// func playGame() {
//     rand.Seed(time.Now().UnixNano()) // Инициализируем рандомизатор

//     target := getRandomNumber(1, 100) // Загадываем число
//     var guess int
//     var err error
//     attempts := 0

//     fmt.Println("Компьютер загадал случайное число от 1 до 100 включительно. Угадайте его!")

//     scanner := bufio.NewScanner(os.Stdin)

//     for {
//         fmt.Print("Ваше предположение (либо, для завершения, введите слово выход): ")
//         input := getUserInput(scanner)
//         input = strings.TrimSpace(input)

//         if input == "выход" {
//             fmt.Println("\nВы вышли из игры.")
//             return
//         }

//         guess, err = strconv.Atoi(input)
//         if err != nil {
//             fmt.Println("Пожалуйста, введите корректное число или слово 'выход'.")
//             continue
//         }

//         if guess < 1 || guess > 100 {
//             fmt.Println("Число вне диапазона. Введите число от 1 до 100.")
//             continue
//         }

//         attempts++

//         if guess < target {
//             fmt.Println("Загаданное число больше.")
//         } else if guess > target {
//             fmt.Println("Загаданное число меньше.")
//         } else {
//             fmt.Printf("Правильно! Вы угадали число с %d попытки.\n", attempts)
//             break
//         }
//     }

//     // Предлагаем сыграть снова
//     fmt.Print("Хотите сыграть еще раз? (если хотите, напишите слово да): ")
//     response := getUserInput(scanner)
//     if strings.ToLower(strings.TrimSpace(response)) == "да" {
//         playGame()
//     } else {
//         fmt.Println("Спасибо за игру! До свидания!")
//     }
// }

// func main() {
//     playGame()
// }

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

var ErrUserEndGame =errors.New("user ehd game")


func generateRandomNumber(min,max int) int{
    return rand.IntN(max-min +1)+ min

}



func getUserInput(min,max int) int{
    var input string
    fmt.Printf("Ваше предположение (либо, для завершения, введите слово выход):")
fmt.Sscanln(&input)


if string.ToLower(input) == "выход" {
    return 0, ErrUserEndGame
}
number, err := strconv.Atoi((input))
if err != nil {
   return 0, fmt.Errorf("invalid input: %w", err) 
}
return number, nil
}

func playGame(){
    min := 1
    max := 100
   randomNumber := generateRandomNumber(min, max)
   fmt.Println("Компьютер загадал случайное число от 1 до 100 включительно. Угадайте его!")

   for {
    input, err := getUserInput()
    if err !=nil{
        fmt.Println(err)
        return
    }
    fmt.Println(input)
   }
}


func main (){
   for{
    playGame()
    var playAgain string
    fmt.Println("Хотите сыграть еще раз? (если хотите, напишите слово да): нет")
    fmt.Scanln(&playAgain)
    if playAgain != "да"{
        fmt.Println("Спасибо за игру! До свидания!")
        break
    }
   }
}