

package main

import (
	"bufio"
	"fmt"
	"os"
)

func PrintReplaced(str string){
    for _, ch := range str {
    if ch == 'у' {
			fmt.Print("а")
		} else {
			fmt.Printf("%c", ch)
		}
	}
	fmt.Println()

}