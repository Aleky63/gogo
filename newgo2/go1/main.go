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
func (y YAZ) StepOnGas() {

	fmt.Println("Я YAZik! 22222!")
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
	fmt.Println("")
	ride(yaz)

}
