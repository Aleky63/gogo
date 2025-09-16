package main

import (
	"fmt"
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
func main14() {
    
    testStr := "Привет, улыбка!"
    
    PrintReplaced(testStr)
}