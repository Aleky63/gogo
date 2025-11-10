package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	transferPoint := make(chan int)

	go func() {
		iterations := 3 + rand.Intn(4)

		for i := 1; i <= iterations; i++ {
			transferPoint <- 10
			time.Sleep(300 * time.Millisecond)
		}
		close(transferPoint)
	}()

	coal := 0
	// for {
	// 	v, ok := <-transferPoint
	// 	if !ok {
	// 		fmt.Println("Oбработал всёё---")
	// 		break
	// 	}
	// 	coal += v
	// 	fmt.Println("coal:", coal)
	// }

	for v := range transferPoint {
		coal += v
		fmt.Println("coal:", coal)
	}

	fmt.Println("Summa coal:", coal)

}
