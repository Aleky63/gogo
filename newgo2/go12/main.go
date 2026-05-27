package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

var (
	likes int = 0
	mtx   sync.RWMutex
)

func setLikes(wg *sync.WaitGroup) {
	defer wg.Done()
	mtx.Lock()
	likes++
	mtx.Unlock()
}

func getLikes(wg *sync.WaitGroup) {
	defer wg.Done()
	mtx.RLock()
	_ = likes
	mtx.RUnlock()
}

func main() {
	red := color.New(color.FgHiRed).SprintFunc()
	wg := &sync.WaitGroup{}

	initNime := time.Now()
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go setLikes(wg)
	}
	for i := 0; i <= 10000; i++ {
		wg.Add(1)
		go getLikes(wg)
	}
	fmt.Println(red("Time :", time.Since(initNime)))
	wg.Wait()
}
