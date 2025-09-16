package main

import (
	"fmt"
	"maps"
)

func main44() {
	res := map[string][]int{
		"Alice":   {1, 2},
		"Bob":     {3, 4},
		"Charlie": {7, 8},
		"Dave":    {5, 1},
	}
	fmt.Println(res)
	RemoveSlicesBySum(res)
	fmt.Println(res)
}



func RemoveSlicesBySum(m map[string][]int) {
	maps.DeleteFunc(m, func(key string, value []int) bool {
   sum := 0
   for _, v := range value {
   sum += v

	}
	return sum > 6
})
}



// func sumSlice(slice []int) int {
// 	sum := 0
// 	for _, num := range slice {
// 		sum += num
// 	}
// 	return sum
// }

// func RemoveSlicesBySum(m map[string][]int) {

// 	for k, slice := range m {
// 		if sumSlice(slice) > 6 {
// 			delete(m, k)
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"maps"
// 	"slices"
// )

// func main() {
// 	m := map[string][]int{
// 		"a": {5, 2},
// 		"b": {15, 12},
// 		"c": {5, 20},
// 		"d": {1, 2},
// 		"f": {66, 2},
// 	}
// 	maps.DeleteFunc(m, func(key string, value []int) bool {
// //   for _, v := range value {
// // 	if v ==5 {
// // 		return true
// // 	}

// //   }
// //   return false

//  return slices.Contains(value, 5)

// 	})
// 	fmt.Print(m)
// }
// https://stepik.org/lesson/1500862/step/1?unit=1520979
