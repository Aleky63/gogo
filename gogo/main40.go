package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main40() {
	text := "Да здравствует ку-ку язык, да здравствует 99 ку-ку goland!"

	  text = strings.Map(func(r rune) rune {
if unicode.IsPunct(r)  {
	return -1
}
	
return unicode.ToLower(r)
	}, text)

	words := strings.Split(text, " ")
	fmt.Println(words)

	count:= make(map[string]int)
	for _, word := range words {
count[word]++

		fmt.Println(word)
	}

		fmt.Println(count)

		for word, v := range count {
	suffix := "раз"
	if v%100 >= 11 && v%100 <= 14 {
		suffix = "раз"
	} else {
		switch v % 10 {
		case 1:
			suffix = "раз"
		case 2, 3, 4:
			suffix = "раза"
		default:
			suffix = "раз"
		}
	}
	fmt.Printf("Слово %q встречается в строке %d %s.\n", word, v, suffix)
  }
}

// https://stepik.org/lesson/1500859/step/1?unit=1520976