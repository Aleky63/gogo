package main

import (
	"fmt"
)


func main ()  {
s1 := []int	{1, 2, 3, 4}
s2 := []int	{11, 22, 33, 44, 555}
		
copied :=copy(s2,s1)
	fmt.Println(s2)
	fmt.Print(copied)

	
 }