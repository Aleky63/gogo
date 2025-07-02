package main

import (
	"fmt"
)
func main(){
	fmt.sumOfDigits()
}


 func sumOfDigits(n int) int{
	if n < 0 {
	n = -n	
	}
	if n== 0{
		return 0
	}
return (n %10) + sumOfDigits(n/10)
 }