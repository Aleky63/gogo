package main

import (
	
	"fmt"
)



func Max(nums []int) (int, error){
	if nums == nil || len (nums) == 0 {
	  return 0, fmt.Errorf("slice is nil or empty")
}

max := nums[0]


for i := range nums {
	if nums[i]>max {
		max = nums[i]
	}
	
}
return max, nil
}

func main() {
	// Тестовые случаи
	testCases := [][]int{
		{1, 3, 2, 8, 5},
		{-10, -5, -20, -1},
		{42},
		{},
		nil,
	}
	
	for i, testCase := range testCases {
		max, err := Max(testCase)
		if err != nil {
			fmt.Printf("Тест %d: Ошибка - %v\n", i+1, err)
		} else {
			fmt.Printf("Тест %d: Максимальный элемент = %d\n", i+1, max)
		}
	}
}

