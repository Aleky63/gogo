package main

import (
	"fmt"

	"github.com/fatih/color"
)

var (
red = color.New(color.FgHiRed).SprintFunc()
blue = color.New(color.FgHiMagenta).SprintFunc()

)

func greet (name string){
	fmt.Println(name,"!!!")
}

type Coordinate struct {
	X,Y int
}

func ( coord *Coordinate) shiftBy(x, y int)  {
	coord.X +=x
	coord.Y +=y
	
} 


func main() {
greet("Tramp")



	fmt.Println (red(555) + blue (" dfdfdfdefd  "))


for i := 0; i <= 10; i++ {
	if i%2 !=0 {
	continue
}
fmt.Println(i)
}

type Sasa struct{
	field string 
	a,b,c int
}
data :=Sasa{}
data =Sasa{field: "Tramp"}

fmt.Println(data)


coord := Coordinate{5,5}
coord.shiftBy(1, 1)
fmt.Println(coord)

}


