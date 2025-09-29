package main

import (
	"fmt"

	"github.com/fatih/color"
)

func createTemperatureAdjuster() (func(change float64) float64, float64) {
	baseTemperature := 90.0

	adjustTemperature := func(change float64) float64 {
		baseTemperature = baseTemperature + change
		return baseTemperature
	}
	return adjustTemperature, baseTemperature

}
func main() {
	red := color.New(color.FgRed).SprintfFunc()

	adjustTemp, origionTemp := createTemperatureAdjuster()

	fmt.Printf("Current temp is %s°C, original was %s°C\n", red("%.2f", adjustTemp(1.5)), red("%.2f", origionTemp))
	fmt.Printf("Current temp is %s°C, original was %s°C\n", red("%.2f", adjustTemp(1.5)), red("%.2f", origionTemp))
	fmt.Printf("Original temp remains: %s°C\n", red("%.2f", origionTemp))

}
