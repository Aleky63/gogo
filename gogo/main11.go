package main

import (
	"errors"
	"fmt"
)

func main11() {
	score, err := getScore()
	if err != nil {
		fmt.Println(err)
				return	
	}
	letterGrate := getLetter11(score)
fmt.Printf("Ваша оценка: %s\n", letterGrate)
}

func getScore11() (int, error){
	var score int
	fmt.Printf("Введите вашу оценку (0-100): ")
	if _, err := fmt.Scan(&score); err != nil {
	
		return 0, fmt.Errorf("input error: please enter the whole scoreber: %w",err)
	}

	if score <0 ||  score > 100 {
return 0, errors.New("the score should be in the range from 0 to 100 inclusive.")
	}

return score, nil
} 

func getLetter11(score int)string{

	switch {		
	case score >= 90 && score <= 100:
		return  "A"
	case score >= 80 && score < 90:
		return  "B"
    case score >= 70 && score < 80:
		return  "C"
    case score >= 60 && score < 70:
		return  "D"
	default: 
			return "F"
	}
}
