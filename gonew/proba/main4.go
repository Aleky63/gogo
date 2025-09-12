package main

import (
	"bufio"
	"fmt"
	"os"
)

func main4() {
	in := bufio.NewScanner(os.Stdin)
	var prev string

	// alreadySeen := make(map[string]bool)

	for in.Scan() {
		txt := in.Text()

		// if _, found := alreadySeen[txt]; found {
		// 	continue
		// }
		// alreadySeen[txt] = true

		//   если отсортировано то можно так >>>>
		if txt == prev {
			continue
		}
		if txt < prev {
			panic("rereeere not not")
		}
		prev = txt

		fmt.Println((txt))
	}
}
