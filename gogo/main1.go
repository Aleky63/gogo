package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main1() {
	fmt.Printf("Введите название товара")
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
if err :=scanner.Err(); err!=nil{
log.Fatalf("Ошибка ввода:%s", err)
}

input :=scanner.Text()
input = strings.ToLower(input)

// if strings.Contains("клавиатура jz9", input){
//    fmt.Println("Клавиатура JZ9:",19200 )

// } else if strings.Contains("наушники n45", input){
// 	 fmt.Println("Наушники N45:",9600 )

// } else if strings.Contains("смартфон s10", input){
// 	 fmt.Println("Смартфон S10:",55000 )
// }else{
// 	fmt.Printf("Товар %q не найден.\n", input)
// }

switch {
case strings.Contains("клавиатура jz9", input):
 fmt.Println("Клавиатура JZ9:",19200 )
case strings.Contains("наушники n45", input):
	fmt.Println("Наушники N45:",9600 )
case strings.Contains("смартфон s10", input):
	 fmt.Println("Смартфон S10:",55000 )
default: fmt.Printf("Товар %q не найден.\n", input)
   }
}