package main

import "fmt"



func filterEven(nums ...int) []int{

var res []int
for _, num := range nums {

	if num%2 == 0 {
		res = append(res, num)
	}
}
return res
}

func main28(){
	
	
 n:= filterEven(2,5,6,14,11,85)


 	fmt.Println("Результат:", n)
}


