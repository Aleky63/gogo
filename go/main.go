package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var ErrUserEndGame = errors.New("user ehd game")

func generateRandomNumber(min,max int) int{
    return rand.Intn(max-min +1)+ min
}

func getUserInput(min, max int) (int,error){
    var input string
    fmt.Printf("Ваше предположение (либо, для завершения, введите слово выход):")
fmt.Scanln(&input)

if strings.ToLower(input) == "выход" {
    return 0, ErrUserEndGame
}
number, err := strconv.Atoi((input))
if err != nil {
   return 0, fmt.Errorf("invalid input: %w", err) 
}

if number < min || number > max {
        return 0, fmt.Errorf("число должно быть в диапазоне от %d до %d", min, max)
    }

return number, nil
}

func playGame() error {
    min := 1
    max := 100
   randomNumber := generateRandomNumber(min, max)
   attempts := 0
   
   


fmt.Println("Компьютер загадал случайное число от 1 до 100 включительно. Угадайте его!")

   for {
    input, err := getUserInput(min,max)
    if err !=nil{
        if err == ErrUserEndGame{
            return  err
} 
fmt.Println("Неккоректное число")
continue
        }    
     attempts++   
    
    if input < randomNumber {
         fmt.Println("Загаданное число больше.")
    } else if input > randomNumber{
fmt.Println("Загаданное число меньше.")
    } else {
       fmt.Printf("Правильно! Вы угадали число с %d попытки.\n", attempts) 
       break
    }
   }
 return nil
}

func printGameEnd(){
   fmt.Println("Спасибо за игру! До свидания!") 
}

func main () {
   for {
    if err := playGame(); err != nil  {
        if err == ErrUserEndGame {
            printGameEnd()
            break   
        }
fmt.Printf("Возникла ошибка во время игры: %v", err)
os.Exit(1)     
    }

    var playAgain string
    fmt.Println("Хотите сыграть еще раз? (если хотите, напишите слово да): нет")
    fmt.Scanln(&playAgain)
    if playAgain != "да"{
    printGameEnd()
        break
    }
   }
  
}