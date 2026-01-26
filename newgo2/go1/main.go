package main

import (
	"fmt"

	"github.com/fatih/color"
)

type Skoda struct{}
type YAZ struct{}

type Auto interface {
	StepOnGas()
	StepOnBrake()
	Bibi()
}

func (b Skoda) StepOnGas() {
	fmt.Println("Я Шкода! 55555!")
}

func (b Skoda) StepOnBrake() {
	magenta := color.New(color.FgHiMagenta).SprintFunc()
	fmt.Println(" ")
	fmt.Println(magenta("Skoda braking erereree"))
}
func (b Skoda) Bibi() {
	green := color.New(color.BgHiGreen).SprintFunc()
	fmt.Println(green(" ----  Я Шкода! bibibibibi!"))
}

func (y YAZ) StepOnGas() {
	fmt.Println("Я YAZik! 22222!")
}
func (y YAZ) StepOnBrake() {
	magenta := color.New(color.FgHiMagenta).SprintFunc()
	fmt.Println(" ")
	fmt.Println(magenta("YAZ braking yyyyyyyyy"))
}
func (b YAZ) Bibi() {
	green := color.New(color.BgHiGreen).SprintFunc()
	fmt.Println(green(" ----  Я Буханка! bibibibibi!"))
}
func ride(auto Auto) {
	red := color.New(color.FgHiRed).SprintFunc()

	fmt.Println(red("Я водила с нижнего тагила"))
	fmt.Println("Я садюсь в авто")
	fmt.Println("Я жму на газ")

	auto.Bibi()
	auto.StepOnGas()
	auto.StepOnBrake()
	

}

func main() {

	sk:= Skoda{}
	yaz := YAZ{}

	ride(sk)
	
	fmt.Println("")
	ride(yaz)

}
