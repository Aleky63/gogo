package main

import (
    "fmt"
)

func main9() {
    result := sumOfDigits(12345)
    fmt.Println("Сумма цифр:", result)
}

func sumOfDigits(n int) int {
    if n < 0 {
        n = -n
    }
    if n == 0 {
        return 0
    }
    return (n % 10) + sumOfDigits(n / 10)
}