package main

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
)

// var number int = 0
var slice []int

var mtx sync.Mutex

func increase(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10000; i++ {
		mtx.Lock()
		slice = append(slice, i)
		mtx.Unlock()
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
	fmt.Println(red(len(slice)))
}
