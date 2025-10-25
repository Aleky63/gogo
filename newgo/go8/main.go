package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("ВВЕДИТЕ КОМАНДУ: ")
		ok := scanner.Scan()
		if !ok {
			fmt.Println("Ошибка ввода")
			return
		}
		text := scanner.Text()
		fmt.Println("text:", text)

		fields := strings.Fields(text)
		if len(fields) == 0 {
			fmt.Println("Вы ничего не ввели")
			return
		}
		pp.Println("слова:", fields)
		fmt.Println("команда:", fields[0])
		cmd := fields[0]

		if cmd == "выйти" {
			fmt.Print("poka")
			return
		}
		if cmd == "добавить" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]
				if i != len(fields)-1 {
					str += " "
				}

			}

			fmt.Println("Вы хотите добавить:", str)
		} else if cmd == "удалить" {
			fmt.Println("Вы хотите удалить")
		} else if cmd == "help" {
			fmt.Println("Команда: добавить(что нужно добавить)")
			fmt.Println("--эта команда позволяет добавлять что-то)")
			fmt.Println("")

		} else {
			fmt.Println("Вы ввели неизвестную команду")
		}
	}
}

// https://www.youtube.com/watch?v=dpvRDJjUJf8
