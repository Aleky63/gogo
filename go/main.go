package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
)

func main() {
slice,err := CreateSlice(10)

if err != nil {
	log.Fatalf("ошибка при создании слайсаЖ %v", err)
}
 fmt.Println(  "сгенерируемый слайс", slice)	



 filtered := FilterSlice(slice)
 fmt.Println("отфильтрованный слайс", filtered)	
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






// func MaxSumWithNegative(numbers []int, k int) []int {

// }

// func SortByParity(numbers []int) []int {

// }

