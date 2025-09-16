package main

import (
	"fmt"
	"sort"
	"strings"
)

func main41() {
	votes1 := map[string]int{
		"Алиса":   15,
		"Боб":     110,
		"Виктор":  288,
	}
	countVotes(votes1)
	 fmt.Println(countVotes(votes1))
}

func countVotes(votes map[string]int ) string {

if   votes == nil || len(votes) == 0  {
		return "Кандидаты потерялись."
	}

	max := 0
	for _, v := range votes {
		if v > max {
     max = v
		}		
	}
	
	if max == 0{
		return "Все голоса похищены!"
	}

var winners []string 
for candidate, count := range votes {
	if count == max{
		winners =append(winners,candidate)
	}
}

sort.Strings(winners)
return strings.Join(winners, ", ")
}

// https://stepik.org/lesson/1500859/step/2?thread=solutions&unit=1520976