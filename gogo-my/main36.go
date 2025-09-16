package main

import (
	"fmt"
	"unicode/utf8"
)

func main36() {
	str := "Зашифруй меня!"
	encodedStr := caesarCode(str, 5, true)
	fmt.Println(encodedStr)

	decodedStr := caesarCode(encodedStr, 5, false)
	fmt.Println(decodedStr)
}
func caesarCode(text string, shift int32, encode bool) string{
	result := make([]rune, 0, utf8.RuneCountInString(text	))
for _, r := range text {

	if encode {
		result = append(result, r+shift)
	}else{
result = append(result, r-shift)
	}	
	// fmt.Println(r)
	
}
return  string(result)
}
// https://stepik.org/lesson/1500852/step/2?unit=1520969