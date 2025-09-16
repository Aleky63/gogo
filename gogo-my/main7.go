

package main
import (
	"fmt"
)

func main7(){
	add := adder(11)
	fmt.Println(add(5))
	fmt.Println(add(5))
	fmt.Println(add(5))
}

func adder(n int) func(x int) int {
    sum := n
    return func(x int) int {
        sum += x
        return sum
    }
}