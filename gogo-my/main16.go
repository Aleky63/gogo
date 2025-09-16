package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func printTable(num int) {
    for i := 1; i <= num; i++ {
        for j := 1; j <= num; j++ {
            if j > 1 {
                fmt.Print("\t")
            }
            fmt.Printf("%d x %d = %d", i, j, i*j)
        }
        fmt.Println()
    }
}

func main16() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    num, _ := strconv.Atoi(scanner.Text())
    printTable(num)
}