package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name   string
	Grades []int
}

func (s Student) AverageGrade() float64 {

	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0
	for _, grade := range s.Grades{
	sum += grade	
	}
	 averageGrade := float64(sum)/float64(len(s.Grades))

 return math.Round( averageGrade*10)/ 10
}

func (s Student) Info() string {
// avg := s.AverageGrade()
return fmt.Sprintf("Студент:%s, средняя оценка: %.1f.", s.Name, s.AverageGrade())


	
}

func main() {
	user := Student {
	Name: "Aenby",
	Grades:[]int{5,5,5,4,5,4,5,4,5,3,5,5,5,4,5,3,5,5},
}
	user2 := Student {
	Name: "Frrrrrr",
	Grades:[]int{5,5,5,3,5,4,5,3,3,3,4},
}
fmt.Println (user)
fmt.Println ( user2)

fmt.Println(user.Info())
fmt.Println(user2.Info())



}

	// https://stepik.org/lesson/1500869/step/3?unit=1520986

