package miner

import (
	"context"
	"fmt"
	"time"
)

func Miner(ctx context.Context, transferPoint chan<- int, n int, power int) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я шахтер номер:", n, "Мой рабочий день закончен!")
			return
		default:
			fmt.Println("Я шахтер номер:", n, "Начал добывать уголь!")
			time.Sleep(1 * time.Second)
			fmt.Println("Я шахтер номер:", n, "Добыл уголь!", power)

			transferPoint <- power
			fmt.Println("Я шахтер номер:", n, "Передал уголь!", power)
		}

	}

}
func foo(ctx context.Context, minerCount int) {
	coalTransferPoint := make(chan int)

	for i := 1; i <= minerCount; i++ {
		go Miner(coalTransferPoint, i, i*10)
	}
}
