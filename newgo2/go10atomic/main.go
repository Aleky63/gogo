package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var num atomic.Int64

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {

		num.Add(1)
	}

}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(11)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	wg.Wait()
	fmt.Println("num:", num.Load())
	fmt.Println("🤣--END--😊")
}
