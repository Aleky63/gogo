package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"slices"
	
)

func main() {
slice,err := CreateSlice(10)

if err != nil {
	log.Fatalf("ошибка при создании слайсаЖ %v", err)
}
 fmt.Println(  "сгенерируемый слайс", slice)	



 filtered := FilterSlice(slice)
 fmt.Println("отфильтрованный слайс", filtered)	


 resMax := MaxSumWithNegative(slice,3)
 if resMax != nil{
 fmt.Println("максимальная сумма слайс", resMax)	
 }else{
	 fmt.Println("нет такого слайса")	
 }
 sortBy := SortByParity(slice)
fmt.Println("отсортированный слайс", sortBy)	
}

func CreateSlice(n int) ([]int, error) {
if n < 0 {
	return nil, errors.New("negative n")
}

s:= make([]int, n)
for i := 0; i< n; i++ {
	s[i] =rand.IntN(21) - 10
}
return s, nil
}

func FilterSlice(numbers []int) []int {

	filtred := []int{}
	
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] &&
			(numbers[i]%2 == 0 ||
				numbers[i]%5 == 0 ||
				numbers[i]%6 == 0 ||
				numbers[i]%9 == 0) {
			filtred = append(filtred, numbers[i])
		}
	
	}
	return filtred
} 

func MaxSumWithNegative(numbers []int, k int) []int {
maxSum := math.MinInt
	var res []int

	for i := 0; i < len(numbers)-k+1; i++ {
		subslice := numbers[i : i+k]
		sum := 0
		hasNegative := false
		for _, num := range subslice {
			sum += num
			if num < 0 {
               hasNegative = true
			 }
		}	
		if hasNegative && sum > maxSum{
			maxSum = sum
			res =  append([]int{},  subslice...)
		}
	}
return res
}

func SortByParity(numbers []int) []int {
	even := []int{}
	odd := []int{}
for _, num := range numbers {
	if num%2 == 0 {
		even = append(even, num)
	}else{
		odd = append(odd, num)
	} 
  }
slices.Sort(even) 
slices.Reverse(even)
slices.Sort(odd)
  return append(even, odd...)
}

