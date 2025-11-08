package main

import (
	"fmt"
	"time"
)

func nine(n int) int {
	fmt.Println("Поход в шахту номер", n, "начался ...")

	time.Sleep(2 * time.Second)

	fmt.Println("Поход в шахту номер", n, "закончился.")
	return 10
}

func main() {
	coal := 0
	initTime := time.Now()
	coal += nine(1)
	coal += nine(2)
	coal += nine(3)

	fmt.Println("Добыли", coal, "угля!")
	fmt.Println("Прошло времени:", time.Since(initTime))
}
