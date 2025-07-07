package main

import "fmt"
func main() {
	nums :=[15]int{1:87, 6:66, 7:2,8:11}
	fmt.Println(nums[6])

	i:=4
	fmt.Println(len(nums))
	if i >=0 && i <=len(nums){
		fmt.Println(nums[i])
	} else {
		fmt.Printf("Значение %d выходит за пределы [от 0 до %d]\n", i, len(nums)-1)
	}
}
