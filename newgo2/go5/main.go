package main

import (
	"fmt"
	"time"
)

func mine(transferPoint chan int, n int) {

	fmt.Println("Поход в шахту номер", n, "начался ...")

	time.Sleep(10 * time.Millisecond)

	fmt.Println("Поход в шахту номер", n, "закончился.")

	transferPoint <- 10
	fmt.Println("Поход №", n, "сдал уголь.")

}

func main() {
	transferPoint := make(chan int, 10)
	coal := 0

	initTime := time.Now()

	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	coal += <-transferPoint
	time.Sleep(1 * time.Second)
	coal += <-transferPoint
	time.Sleep(1 * time.Second)
	coal += <-transferPoint
	time.Sleep(1 * time.Second)

	fmt.Println("Добыли", coal, "угля!")
	fmt.Println("Прошло времени:", time.Since(initTime))
}
