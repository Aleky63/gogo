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


type Direction byte

const (
	North Direction =iota
	East
	South
	West
)


func(d Direction) String() string {
switch d {

case North:
	return "North"
case South:
	return "South"
case East:
	return "East"
case West:
	return "West"
default:
	return "other direction"
	
   }
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
coord.shiftBy(1, 1)
coord.shiftBy(1, 1)
fmt.Println(coord)


x := East.String() 
fmt.Println(x)

}


