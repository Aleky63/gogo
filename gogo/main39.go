package main

import (
	
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
    "banana":     2,
    "apple":      1,
    "grapefruit": 3,
    "cherry":     1,
  }
  invertedMap := invertMap(m)
  printMap(invertedMap)
}

func invertMap(inputMap map[string]int) map[int][]string {
 inverted := make(map[int][]string)
 for key, value := range inputMap {
        inverted[value] = append(inverted[value], key)
    }
	for _, strings := range inverted {
		sort.Strings(strings)
	}
    return inverted
}

func printMap(m map[int][]string) {
	
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	
	fmt.Println("{")
		
	for _, key := range keys {
		values := m[key]
		fmt.Printf("  %d: [", key)
		
		
		for i, value := range values {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("\"%s\"", value)
		}
		
		fmt.Println("],")
	}
	
	fmt.Println("}")
}