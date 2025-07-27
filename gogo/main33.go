// https://stepik.org/lesson/1500848/step/6?unit=1520965

// package main

// import (
// 	"fmt"
// 	"slices"
// )

// func main(){
// 	slice := []int{1,2,3,4,5}
// 	fmt.Println(slice)

// 	slice2 :=  slice[:len(slice)-1]
// 	fmt.Println(slice2)

// 	slice3 :=  slice[1:]
// 	fmt.Println(slice3)

// 	excess := 2
// 	slice4 :=  slice[excess+1:]
// 	slice5 :=  slice[:excess]
// 	slice6 := append(slice5, slice4...)
//     fmt.Println(slice6," ", cap(slice6)," ", len(slice6))

// 	slice6 = slices.Clip(slice6)
//  fmt.Println(slice6," ", cap(slice6)," ", len(slice6))

// slice2 =slice2[:len(slice2):len(slice2)]
// fmt.Println(slice2," ", cap(slice2)," ", len(slice2))

// newSlice := make([]int, len(slice6))
// copy(newSlice, slice6)
// 	 fmt.Println(newSlice," ", cap(newSlice)," ", len(newSlice))

// 	newSlice = slices.Clip(newSlice)
// 		 fmt.Println(newSlice," ", cap(newSlice)," ", len(newSlice))

// }

package main

import (
	"fmt"
	"slices"
)


func DeletingFromSlice(slice[]int) []int{
result :=make([]int, len(slice))
 copy(result, slice)

 flagLast := false
    flagIndex2 := false

   if len(result) > 0 {
 lastElement := result[len(result)-1]

	if lastElement >= 11 {
     result = result[:len(result)-1]
	 flagLast = true
    }
}
 
if len(result) > 2 && cap(slice) >= 4 {
    
        result = append(result[:2], result[3:]...)
		flagIndex2 = true 
}

if flagLast && flagIndex2 && len(result)> 0{
	result = result[1:]
}
result = slices.Clip(result)
return result
	}

func main33() {

	slice:=[]int{1,2,3,4,5,16,19,66}
	newSlice := DeletingFromSlice(slice)
	  fmt.Println(newSlice)
	  fmt.Println (cap(newSlice))

}
