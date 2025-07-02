package main

import (
    "fmt"
    "io"
    "os"
    "regexp"
    "strconv"
    "strings"

    "math/rand/v2"
)

// rollDice симулирует броски двух кубиков до получения целевой суммы
func rollDice(targetSum int) {
    rollCount := 0
    
    for {
        rollCount++
        
        // Генерируем случайные значения для двух кубиков (от 1 до 6)
        dice1 := rand.IntN(6) + 1
        dice2 := rand.IntN(6) + 1
        sum := dice1 + dice2
        
        if sum == targetSum {
            // Целевая сумма достигнута
            var rollWord string
            var verbForm string
            lastDigit := rollCount % 10
            lastTwoDigits := rollCount % 100
            
            if lastTwoDigits >= 11 && lastTwoDigits <= 14 {
                rollWord = "бросков"
                verbForm = "потребовалось"
            } else if lastDigit == 1 {
                rollWord = "бросок"
                verbForm = "потребовался"
            } else if lastDigit >= 2 && lastDigit <= 4 {
                rollWord = "броска"
                verbForm = "потребовалось"
            } else {
                rollWord = "бросков"
                verbForm = "потребовалось"
            }
            
            fmt.Printf("Выпало %d и %d, в сумме %d, на это %s %d %s.\n", 
                dice1, dice2, sum, verbForm, rollCount, rollWord)
            break
        } else {
            // Целевая сумма не достигнuta
            fmt.Printf("Выпало %d и %d, в сумме %d, бросаем еще раз.\n", 
                dice1, dice2, sum)
        }
    }
}

func main() {
    content, err := io.ReadAll(os.Stdin)
    if err != nil {
        panic(err)
    }

    ioContent := strings.TrimSpace(string(content))
    if ioContent != "" {
        targetSum, _ := strconv.Atoi(ioContent)
        rollDice(targetSum)
    }
}