package main

import (
	"fmt"
)

type Day int

const (
_Day = iota
Monday
Tuesday
Wednesday
Thursday
Friday
Saturday
Sunday
)

func isWeekend(d Day) bool {
	return d == 6 || d == 7
}

func main8(){
  fmt.Println(isWeekend(Saturday))
}

