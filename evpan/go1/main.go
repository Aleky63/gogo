package main

import "fmt"

// Интерфейс для дома
type House interface {
	HasWalls() bool
	HasFloor() bool
	HasRoof() bool
	HasDoors() bool
}

// Реализация интерфейса
type Dacha struct {
	Walls bool
	Floor bool
	Roof  bool
	Doors bool
}

func (d Dacha) HasWalls() bool {
	return d.Walls
}

func (d Dacha) HasFloor() bool {
	return d.Floor
}

func (d Dacha) HasRoof() bool {
	return d.Roof
}

func (d Dacha) HasDoors() bool {
	return d.Doors
}

func IsHabitable(h House) bool {
	return h.HasWalls() && h.HasFloor() && h.HasRoof() && h.HasDoors()
}

func main() {
	dacha := Dacha{Walls: true, Floor: true, Roof: true, Doors: true}
	fmt.Println("Дом готов к заселению:", dacha.HasWalls() && dacha.HasFloor() && dacha.HasRoof() && dacha.HasDoors())
}
