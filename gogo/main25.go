package main

import "fmt"

func main(){



slice := make([]int, 5, 5)
fmt.Printf("slice: %v, len: %d, cap: %d\n", slice, len(slice),cap(slice))

slice2 := make([]int, 5, 10)
fmt.Printf("slice2: %v, len: %d, cap: %d\n", slice2, len(slice2),cap(slice2))

slice4 := make([]int, 0)
slice5 := make([]int, 0, 100)

for i := 1; i <= 10; i++ {
	slice4 = append(slice4, i)
	slice5 = append(slice5, i)
	fmt.Printf("slice4 cap: %d, slice5 cap: %d\n", cap(slice4), cap(slice5))
}

nums :=[]int{1,2,3,5,8,6}
numbers :=[]int{1,2,3,5,8,6}
  for _, number := range numbers {
    number *= 10
	fmt.Printf("number: %d\n", number)
}
for i := range nums {
	nums[i] *= 10
	
}
fmt.Printf("numbers: %d", numbers)

fmt.Printf("nums: %d", nums)

}