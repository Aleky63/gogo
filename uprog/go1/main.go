package main

import (
	"fmt"

	"github.com/fatih/color"
)



func greet (name string){
	fmt.Println(name,"!!!")
}


func main() {
greet("Tramp")

var (
red = color.New(color.FgHiRed).SprintFunc()
blue = color.New(color.FgHiBlue).SprintFunc()

)

	fmt.Println (red(555) + blue (" dfdfdfdefd  "))


for i := 0; i <= 10; i++ {
	if i%2 !=0 {
	continue
}
fmt.Println(i)
}

type S struct{
	field string 
	a,b,c int
}
data :=S{}
data =S{field: "Tramp"}

fmt.Println(data)
}


