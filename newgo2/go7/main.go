package main

import (
	"fmt"
	"time"
)

type Message struct {
	Autor string
	Text  string
}

func main() {
	messagechan1 := make(chan Message)
	messagechan2 := make(chan Message)

	go func() {

		for {
			messagechan1 <- Message{
				Autor: "Tramp",
				Text:  "Skazka",
			}
			time.Sleep((33 * time.Millisecond))
		}
	}()
	go func() {

		for {
			messagechan2 <- Message{
				Autor: "Baiden",
				Text:  "Pokoy",
			}
			time.Sleep((100 * time.Millisecond))
		}
	}()

	start := time.Now()

	for time.Since(start) < 3*time.Second {
		select {
		case msd1 := <-messagechan1:
			fmt.Println("Я получил сообщение от:", msd1.Autor, "текс сообщения:", msd1.Text)
		case msd2 := <-messagechan2:
			fmt.Println("Я получил сообщение от:", msd2.Autor, "текс сообщения:", msd2.Text)

		}

	}
	fmt.Println("\n⏹️  Время вышло — программа остановлена.")
}
