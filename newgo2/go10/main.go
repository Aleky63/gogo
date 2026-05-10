package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/fatih/color"
)

// var number int = 0
var number atomic.Int64

func increase(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10000; i++ {
		number.Add(1)
	}
}

func main() {
	red := color.New(color.FgHiRed).SprintFunc()

	wg := &sync.WaitGroup{}
	wg.Add(10)

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
	fmt.Println(red(number.Load()))
}
