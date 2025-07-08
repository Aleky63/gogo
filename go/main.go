package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [7]int{1, 9, 6, 3, 1, 9, 8}
	fmt.Println("Cекретный код:", func() string {
	
		return generateCode(arr)
	}())

}

func generateCode(arr [7]int) string {
	min := arr[0]
    max := arr[0]
     for _, v := range arr {
        if v < min {
            min = v
        }
        if v > max {
            max = v
        }
    }

    var middle string
    for _, v := range arr {
        prefix := "O"
        if v%2 == 0 {
            prefix = "E+/+E"
        }
        middle += prefix + strconv.Itoa(v) 
    }

    
    return strconv.Itoa(min)  + middle +  strconv.Itoa(max)
}
