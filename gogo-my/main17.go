package main

import (
	"fmt"
)

func printDiamond(n int) {
    fmt.Println("Мой бриллиант:")


    for i := 0; i < n; i++ {
        for j := 0; j < n-i-1; j++ {
            fmt.Print(" ")
        }

        fmt.Print("#")


        middleSpaces := 2*i - 1
        if middleSpaces > 0 {
            for j := 0; j < middleSpaces; j++ {
                fmt.Print(" ")
            }
            fmt.Print("#")
        }

        fmt.Println()
    }


    for i := n - 2; i >= 0; i-- {
        for j := 0; j < n-i-1; j++ {
            fmt.Print(" ")
        }

        fmt.Print("#")

        middleSpaces := 2*i - 1
        if middleSpaces > 0 {
            for j := 0; j < middleSpaces; j++ {
                fmt.Print(" ")
            }
            fmt.Print("#")
        }

        fmt.Println()
    }
}

func main17() {
    printDiamond(9)
}