package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		i := 1
		for {
			intCh <- i
			i++

			time.Sleep(1 * time.Second)
		}

	}()
	go func() {
		i := 1
		for {
			strCh <- "hirere55earrfxXSrWW--" + strconv.Itoa((i))
			i++
		}

	}()

	for {
		select {

		case num := <-intCh:
			fmt.Println("intCh:", num)
		case str := <-strCh:
			fmt.Println("strCh:", str)
		default:
			fmt.Println("никакой канал не готов")
		}

	}
}
