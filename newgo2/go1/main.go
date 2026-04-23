package main

import (
	"fmt"

	"github.com/fatih/color"
)

type Auto interface {
	StepOnGas()
	StepOnBrake()
}
type Honker interface {
	Bibi()
}
type (
	Skoda struct{}
	YAZ   struct{}
	BMW   struct{}
)

func (b BMW) StepOnGas() {
	fmt.Println("Я BMW! 55555!")
}

func (b BMW) StepOnBrake() {
	fmt.Println("Я BMW! торможу!")
}

func (b BMW) Bibi() {
	fmt.Println("Я BMW! бибибика!")
}

func (s Skoda) StepOnGas() {
	fmt.Println("Я Шкода! 33333!")
}

func (s Skoda) StepOnBrake() {
	fmt.Println("Я Шкода! торможу!")
}

func (y YAZ) StepOnGas() {
	fmt.Println("Я YAZ!  😍YYYYYYYYYYYYYYYYYYYYYYYY!")
}

func (y YAZ) StepOnBrake() {
	fmt.Println("Я YAZ!  😍торможжжжу!")
}

func ride(auto Auto) {
	red := color.New(color.FgHiRed).SprintFunc()

	fmt.Println(red("Я водила с нижнего тагила"))
	fmt.Println("Я садюсь в авто")
	fmt.Println("Я жму на газ")
	auto.StepOnGas()
	auto.StepOnBrake()

	if h, ok := auto.(Honker); ok {
		h.Bibi()
	}
}

func main() {
	bmw := BMW{}
	skoda := Skoda{}
	yaz := YAZ{}

	ride(bmw)
	ride(skoda)
	ride(yaz)
}
