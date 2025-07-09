package main

import (
    "fmt"
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
}

func generateCompliment(name string) string {
    number := rand.Intn(3) // Генерируем число от 0 до 2
    var res string

    switch number {
    case 0:
        res = "Ты великолепен,"
    case 1:
        res = "У тебя потрясающая улыбка,"
    case 2:
        res = "Ты вдохновляешь,"
    }

    return fmt.Sprintf("%s %s!", res, name)
}

func main5() {
    fmt.Println(generateCompliment("Анна"))
}