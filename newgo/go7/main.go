package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/k0kubun/pp"
)

func main() {
	red := color.New(color.FgHiRed).SprintFunc()
	blue := color.New(color.FgHiBlue).SprintFunc()

	weather := map[int]int{
		11: +3,
		12: +13,
		13: +8,
		14: +5,
		15: +1,
		16: -6,
	}
	weather[20] = +30

	for k, v := range weather {
		fmt.Println(k, red(v))
		weather[k] += 6

	}
	fmt.Println(blue(weather))

	pp.Println(weather)

	crim := map[string]bool{
		"vasy": true,
		"pety": true,
		"alex": false,
	}
	c, ok := crim["alex"]
	if !ok {
		fmt.Println("person not found in database")
		return
	}
	fmt.Println("person found in database ")

	if c {
		fmt.Println("A man is judged")
	} else {
		fmt.Println("A man is not judged")
	}

}
