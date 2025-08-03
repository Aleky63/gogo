package main

import "fmt"

func main38() {
	fmt.Println(CountMaxFrequency([]int{2,5,5,5,6,9,2,8,3,4}))
}

func CountMaxFrequency(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	res := make(map[int]int)

	for _, num := range nums {
		res[num]++

	}

	maxRes := 0
	for _, count := range res {
		if count > maxRes {
			maxRes = count
		}
	}
	return maxRes
}