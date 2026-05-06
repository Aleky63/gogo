package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Message struct {
	Author string
	Text   string
}

func main() {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Println(blue("hello"))

	messadeChan1 := make(chan Message)
	messadeChan2 := make(chan Message)

	go func() {
		for {
			messadeChan1 <- Message{
				Author: "Tramp",
				Text:   "Redhead",
			}

			time.Sleep(10 * time.Second)

		}
	}()

	go func() {
		for {
			messadeChan2 <- Message{
				Author: "Putin",
				Text:   "Mole",
			}

			time.Sleep(1 * time.Second)

		}
	}()

	for {
		select {
		case msg1 := <-messadeChan1:
			fmt.Println(msg1.Author, msg1.Text)
		case msg2 := <-messadeChan2:
			fmt.Println(msg2.Author, msg2.Text)
		}
	}
}
