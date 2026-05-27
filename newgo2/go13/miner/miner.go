package miner

import (
	"context"
	"fmt"
	"time"
)

func Miner(transferPoint chan<- int, n int, power int) {
	for {
		select {
		case <-ctx.Done():

			fmt.Println("I'm a miner number:", n, "the working day is over")
			return
		default:
			fmt.Println("I'm a miner number:", n, "started mining coal")
			time.Sleep(1 * time.Second)

			fmt.Println("I'm a miner number:", n, "mined coal", power)
			transferPoint <- power
			fmt.Println("I'm a miner number:", n, "donated coal", power)
		}
	}
}

func foo(ctx context.Context, minerCount int) chan int {
	coalTransferPoint := make(chan int)
	for i := 0; i <= minerCount; i++ {
		go Miner(ctx, coalTransferPoint, i, i*11)
	}
	return coalTransferPoint
}
