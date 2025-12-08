package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(wg *sync.WaitGroup, text string) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		fmt.Println("Я почтальон, я отнес газету", text, "в", i, "раз")
		time.Sleep(500 * time.Millisecond)
	}

}

func main() {

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go postman(wg, "NEWS")
	wg.Add(1)
	go postman(wg, "The Times")
	wg.Add(1)
	go postman(wg, "AUTO")
	wg.Add(1)
	go postman(wg, "🎈ORORORO")
	wg.Wait()

	fmt.Println("😊😊----END-----🤣🤣")
}
