package main

import (
	"fmt"
)


func PlayWithSlice(slice[]int) []int {
	   
   clon:= make([]int, len(slice))
   	copy(clon,slice)

index := -1

for i := len(clon) - 1; i >= 0; i-- {
	if clon[i] > 10 {
		index = i
		break
	}
}

if index != -1 {
	clon = append(clon, 0)
	copy (clon[index+2:],  clon[index+1:])
	clon[index+1] = 100

}

sum :=0
for _, value := range clon {
	sum +=value
	
}
if sum >100{
	clon = append (clon, 500)
}

countOdd :=0
countEven :=0

for _,value  := range slice {

	if value%2 == 0{
		countEven++
	} else {
		countOdd++
  }
}	
	if countEven > countOdd {
	clon = append ([]int{1000}, clon...)
	}

return clon
}

func main32(){
	res := PlayWithSlice([]int{2, 105, 4, 11, 22, 66, 9})
	fmt.Println(res)

}

// https://stepik.org/lesson/1500847/step/4?unit=1520964