package main

import (
	"fmt"
	"maps"
)

func main43() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"b": 10,
		"c": 5,
		"d": 7,
	}

	merged := mergeMaps(m1, m2)

	fmt.Println(merged)
}
// func mergeMaps(m1, m2 map[string]int) map[string]int {

// 	res := make(map[string]int)

// 	for key, v := range m1 {
// 		res[key] = v

// 	}
// 	for key, v := range m2 {
// 		res[key] += v

// 	}
// 	return res
// }
func mergeMaps(m1, m2 map[string]int) map[string]int {
    res := maps.Clone(m1);
    for key, val := range m2 {
        res[key] += val;
    }
    return res;
}