package main

import (
	"fmt"
)

func main() {
	s := []int{5, 2, 3, 5, 69, 63, 258}

	fmt.Print(notUnique(s))
}

func notUnique(s []int) bool {
	m := make(map[int]struct{})
	for i := 0; i < len(s); i++ {
		val := s[i]
		if _, ok := m[val]; ok {
			return true
		}
		m[val] = struct{}{}
	}
	return false
}
