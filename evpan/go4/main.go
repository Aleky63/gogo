package main

import (
	"fmt"
	"time"
)

// doSomeWork имитирует работу сотрудника,
// который берётся за несколько задач (в цикле).
func doSomeWork(workerID int) {
	for task := 1; task <= 3; task++ {
		fmt.Printf("Работник %d выполняет задачу %d\n", workerID, task)
		time.Sleep(time.Second) // Имитация трудоёмкости задачи
	}
	fmt.Printf("Работник %d закончил все свои задачи\n", workerID)
}

// main это тоже горутина
func main() {
	// Нанимаем двух сотрудников (запускаем горутины)
	go doSomeWork(1)
	go doSomeWork(2)
	go doSomeWork(3)

	fmt.Println("Сотрудники наняты, офис работает...")

	// Даём им время выполнить задачи
	time.Sleep(8 * time.Second)

	fmt.Println("Офис закрывается, все сотрудники завершают работу.")
}

// https://pantela.notion.site/Go-EvgenyPantela-2622eb3830de43639fbb7a0e2b7c2757
