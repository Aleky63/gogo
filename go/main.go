package main

import "fmt"



func SumSlices(s1, s2 []int) []int{
min :=len(s1)
if len(s2) < min {
	min = len(s2)
}

res :=make([]int, min)

for i := range min {
	res[i] = s1[i] + s2[i]
}

return res
}

func main(){
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{10, 20, 30, 40}
	
 sums:= SumSlices(s1 ,s2)
 	fmt.Println("Результат:", sums)
}


