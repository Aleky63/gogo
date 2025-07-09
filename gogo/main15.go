package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
)

var secret = rand.IntN(100) + 1
var attempts int


func guess(num int) (int, error) {
	attempts++
	if attempts > 6 {
		return 0, errors.New("too many attempts")
	}
	if num < secret {
		return 1, nil
	} else if num > secret {
		return -1, nil
	}
	return 0, nil
}


func play() int {
	low := 1
	high := 100

	for i := 0; i < 6; i++ {
		mid := (low + high) / 2
		res, _ := guess(mid)

		if res == 0 {
			return mid
		} else if res == 1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	
	// fallback, по условиям этого не должно быть
	return low
}
func main15() {
	result := play()
	if result == -1 {
		fmt.Println("Не удалось угадать число за 6 попыток")
		os.Exit(1)
	}
	fmt.Println(result)
}
