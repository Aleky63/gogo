package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func main() {
	red := color.New(color.FgRed).SprintFunc()
	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		i := 1
		for {
			intCh <- i
			i++
			time.Sleep(3000 * time.Millisecond)
		}
	}()
	go func() {
		i := 1
		for {
			strCh <- "hi" + strconv.Itoa(i)
			i++
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	fmt.Println(red("hello"))

	for i := 1; i < 10; i++ {
		select {
		case num := <-intCh:
			fmt.Println("intCh", num)
		case str := <-strCh:
			fmt.Println("strCh", str)

		}
	}
}
