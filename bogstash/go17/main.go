package main

import "fmt"

func createTemperatureAdjuster() (func(change float64) float64, float64) {
	baseTemperature := 90.0

	adjustTemperature := func(change float64) float64 {
		baseTemperature = baseTemperature + change
		return baseTemperature
	}
	return adjustTemperature, baseTemperature

}
func main() {

	adjustTemp, origionTemp := createTemperatureAdjuster()
	fmt.Printf("AAAAA is  %.2f grad C\n SSSSS is  %.2f\n ", adjustTemp(1.5), origionTemp)
	fmt.Printf("AAAAA is  %.2f grad C\n SSSSS is  %.2f\n ", adjustTemp(1.5), origionTemp)

}
