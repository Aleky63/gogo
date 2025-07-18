package main

import (
	"fmt"
)


 func replaceEvenOnEvenIndices(slice [][]int) [][]int{
	result := make([][]int, len(slice))

for i, innerSlice := range slice {
	result[i] = make([]int, len(innerSlice))
   for j , v := range innerSlice {
	if j%2 == 0 && v%2 == 0 { 
		result [i][j] = 0
	  }	else {
		result[i][j] = v
			}
	}	
}
return result
 }

func main ()  {
s := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
}
res := replaceEvenOnEvenIndices(s)
	

	fmt.Println(res)
 }