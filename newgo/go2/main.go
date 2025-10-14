package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

func main() {
	var score int
	blue := color.New(color.FgHiBlue).SprintfFunc()
	red := color.New(color.FgHiRed).SprintfFunc()
	score = 11
	if score > 10 && score < 20 {
		fmt.Println(blue("AAAAAAAAAAAAA"))
	} else {
		fmt.Println("ooooooo")
	}

	score2 := false
	if !score2 {
		fmt.Println("DDDDDDDDDDDс❤️")
	} else {
		fmt.Println("iiiiiiiiiiii")
	}

	for i := 1; i < 7; i++ {
		if i%2 == 0 {
			fmt.Println(blue("%v😂", i))
			time.Sleep(1000 * time.Millisecond)
		} else {
			fmt.Println(red("%v A", i))
		}

		fmt.Println("y", i)
		if rand.Intn(4) == 1 {
			fmt.Println("🛠️🛠️🛠️🛠️")
			break
		}
	}

}
