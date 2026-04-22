package main

import (
	"fmt"

	"github.com/fatih/color"
)

type Auto interface {
	StepOnGas()
	// StepOnBrake()
	// Bibi()
}
type (
	Skoda struct{}
	YAZ   struct{}
	BMW   struct{}
)

func (b BMW) StepOnGas() {
	fmt.Println("Я BMW! 55555!")
}

func (s Skoda) StepOnGas() {
	fmt.Println("Я Шкода! 33333!")
}

// func (s YAZ) StepOnGas() {
// 	fmt.Println("Я YAZ!  😍YYYYYYYYYYYYYYYYYYYYYYYY!")
// }

func ride(auto Auto) {
	red := color.New(color.FgHiRed).SprintFunc()

	fmt.Println(red("Я водила с нижнего тагила"))
	fmt.Println("Я садюсь в авто")
	fmt.Println("Я жму на газ")
	auto.StepOnGas()
}

func main() {
	bmw := BMW{}
	skoda := Skoda{}

	// skoda.StepOnGas()
	ride(bmw)
	ride(skoda)
}
