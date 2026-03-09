package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	KickBall()
}

type FootballPlayer struct {
	Power   int
	Stamina int
}

func (f FootballPlayer) KickBall() {
	shot := f.Power + f.Stamina

	fmt.Printf("Я бью по мячу: %d\n", shot)
}

type Messi struct {
	Power   int
	Stamina int
	Height  float64
}

func (m Messi) KickBall() {
	shot := m.Power + m.Stamina
	shot *= 5
	fmt.Printf("Messi ростом: %.2f бьет по мячу: %d", m.Height, shot)
}

func main() {
	players := make([]Player, 11)

	for p := range players {
		players[p] = FootballPlayer{
			Power:   rand.Intn(10),
			Stamina: rand.Intn(10),
		}

		if p == 10 {
			players[p] = Messi{
				Power:   rand.Intn(10),
				Stamina: rand.Intn(10),
				Height:  1.65,
			}
		}

	}

	for p := range players {
		players[p].KickBall()
	}
}
