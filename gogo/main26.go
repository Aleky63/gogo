package main

import "fmt"



func printMagic(nums[]int){
n:=len(nums)

if n == 0 {
		fmt.Println("[]")
		return
}
	result := make([]int, n)


for i := range n {
		multiply := 1
		for j := range n {
			if i != j {
				multiply *= nums[j]
			}
		}
		result[i] = multiply
	}

	fmt.Print("[")
	for i, val := range result {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func main26(){
		nums := []int{1,2,3,4,5}
	printMagic(nums)
}