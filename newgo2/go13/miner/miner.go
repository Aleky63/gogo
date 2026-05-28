package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n int,
	power int,
) {
	defer wg.Done()
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

func Foo(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}
	for i := 0; i <= minerCount; i++ {
		wg.Add(1)
		go Miner(ctx, coalTransferPoint, i, i*11)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
