package main

import (
	"fmt"
)

func main (){
	slice := []int {1,2,4,5}
		
index:=2
value:=3

before := slice[:index]
after:= append([]int{value}, slice[index:]...)        

fmt.Println(before)
fmt.Println(after)

 slice = append(before, after...)
 fmt.Println(slice)

}




