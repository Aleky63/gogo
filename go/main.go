package main

import (
	"fmt"
	"sort"
)


func sortInner(slice []int) {
	sort.Slice(slice, func(i, j int) bool {
		
		if slice[i] == 0 {
			return true
		}
		if slice[j] == 0 {
			return false
		}


		isEvenI := slice[i]%2 == 0
		isEvenJ := slice[j]%2 == 0
		if isEvenI != isEvenJ {
			return isEvenI
		}


		return slice[i] > slice[j]
	})
}


func magicSort(data [][]int) {
	
	for _, inner := range data {
		sortInner(inner)
	}

	
	sort.SliceStable(data, func(i, j int) bool {
		sum := func(s []int) int {
			total := 0
			for _, v := range s {
				total += v
			}
			return total
		}
		return sum(data[i]) < sum(data[j])
	})
}


func main() {
	data := [][]int{
		{5, 2, 0, 3, 8, 80 ,11},
		{1, 7, 3,22,5,66,1025,33 },
		{4, 4, 4, 4}, 
		{9, 0, 2},
	}

	magicSort(data)

	for _, row := range data {
		fmt.Println(row)
	}
}
