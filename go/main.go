package main

import (
	"fmt"
	"strings"
)



// func SumNeighbors(arr [10]int)(newArr [10]int) { 

// newArr = [10]int{}
// newArr[0],newArr[9] = arr[1], arr[8]

// for i := 1; i < len(arr)-1; i++ {
// 	newArr[i] = arr[i-1] + arr[i+1]
// }
// return 
// }

func SumNeighbors(arr [10]int)(newArr [10]int) { 

for i := range arr {
	switch{
	case i == 0:
		newArr[0] = arr[1]
	case i == 9:
		newArr[9] = arr[8]
	 default:
			newArr[i] = arr[i-1] + arr[i+1]
	}
  }
return
}


