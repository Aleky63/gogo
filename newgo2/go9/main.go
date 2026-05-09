package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

func postman(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		fmt.Println("I am postman, I took the newspaper", text, "for", i, "time.")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("hello")

	wg.Add(3)
	go postman("PRABDA", &wg)
	go postman("CRIVDA", &wg)
	go postman("LYLYLYLY", &wg)
	wg.Wait()

	red := color.New(color.FgHiRed).SprintFunc()
	fmt.Println(red("All postmen finished!"))
}
