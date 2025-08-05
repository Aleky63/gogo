// package main

// import (
// 	"fmt"
// 	"maps"
// )

// func main() {
// 	m1 := map[string][]int{
// 		"a": {5, 2},
// 		"b": {50, 20},
// 	}
// 	m2 := map[string][]int{
// 		"a": {5, 2},   // тот же порядок
// 		"b": {20, 50}, // другой порядок
// 	}

// 	res := maps.EqualFunc(m1, m2, func(v1 []int, v2 []int) bool {
// 		if len(v1) != len(v2) {
// 			return false
// 		}
// 		sum1, sum2 := 0, 0
// 		for i := range v1 {
// 			if v1[i] != v2[i] {
// 				return false
// 			}
// 			sum1 += v1[i]
// 			sum2 += v2[i]
// 		}
// 		return sum1 == sum2
// 	})

// 	fmt.Println(res)
// }

package main

import "fmt"

func main() {
	map1 := map[string][]int{
		"a": {1, 3, 5},
		"b": {10, 2},
	}
	map2 := map[string][]int{
		"a": {5, 0},
		"b": {10},
	}
	fmt.Println(CompareMaxValues(map1, map2))
}

func CompareMaxValues(map1, map2 map[string][]int) bool {
	if len(map1) != len(map2)  {
		return false
	}

for key, slice1 := range map1 {
		slice2, ok := map2[key]
		if !ok {
			return false
		}

if len(slice1) == 0 && len(slice2) == 0 {
			continue 
		}
		if len(slice1) == 0 || len(slice2) == 0 {
			return false
		}

		max1 := findMax(slice1)
		max2 := findMax(slice2)

		if max1 != max2 {
			return false
		}
	}

	return true

}
func findMax(slice []int) int {
	if len(slice) == 0 {
		return 0 
	}

	max := slice[0]
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	return max
}