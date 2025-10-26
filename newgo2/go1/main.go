package main

import (
	"fmt"

	"github.com/fatih/color"
)

type BMW struct{}
type YAZ struct{}

type Auto interface {
	StepOnGas()
	StepOnBrake()
}

func (b BMW) StepOnGas() {
	fmt.Println("Я Шкода! 55555!")
}

func (b BMW) StepOnBrake() {
	magenta := color.New(color.FgHiMagenta).SprintFunc()
	fmt.Println(" ")
	fmt.Println(magenta("BMW braking erereree"))
}
func (b BMW) Bibi() {
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

func ride(auto Auto) {
	red := color.New(color.FgHiRed).SprintFunc()

	fmt.Println(red("Я водила с нижнего тагила"))
	fmt.Println("Я садюсь в авто")
	fmt.Println("Я жму на газ")
	auto.StepOnGas()
	auto.StepOnBrake()

}

func main() {

	bmw := BMW{}
	yaz := YAZ{}

	ride(bmw)
	bmw.Bibi()
	fmt.Println("")
	ride(yaz)

}
